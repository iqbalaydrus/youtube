#include <chrono>
#include <cmath>
#include <cstdint>
#include <thread>
#include <iostream>
#include <fstream>
#include <unordered_map>
#include <filesystem>
#include <boost/iostreams/device/mapped_file.hpp>
#include <boost/iostreams/stream.hpp>
#include <boost/utility/string_ref.hpp>
#include <boost/algorithm/string.hpp>

struct locationData {
    uint64_t count = 0;
    double total = 0;
    double temp_min = 99;
    double temp_max = -99;
};

namespace io = boost::iostreams;
const std::string filename = "measurements.txt";

struct Thread {
    std::thread thread;
    long start{};
    long end{};
};

uint64_t find_from_middle(boost::string_ref &s) {
    auto start = (s.begin() + ((s.end() - s.begin()) / 2));
    auto i = 0;
    auto sign = true;
    while (true) {
        if (start > s.end() || start < s.begin()) {
            throw std::invalid_argument("no delimiter found");
        }
        if (*start == ';') {
            break;
        }
        if (sign) {
            start += i;
            sign = false;
        } else {
            start -= i;
            sign = true;
        }
        ++i;
    }
    return start - s.begin();
}

void process_line(Thread &t) {
    std::string line;
    const char delimiter = ';';
    io::stream_buffer<io::mapped_file_source> file(filename);
    std::istream in(&file);
    if (t.start == 0) {
        in.seekg(t.start);
    } else {
        in.seekg(t.start-1);
        char c;
        in.get(c);
        if (c != '\n') {
            std::getline(in, line);
        }
    }
    auto i = 0;
    while (std::getline(in, line)) {
        auto line_ref = boost::string_ref{line};
        uint64_t pos = find_from_middle(line_ref);
        boost::string_ref location = line_ref.substr(0, pos);
        boost::string_ref temperature_str = line_ref.substr(pos+1, line.length());
        auto temperature_num = std::atof(temperature_str.begin());
        if (in.tellg() >= t.end) {
            break;
        }
    }
    file.close();
}

int main_mmap() {
    auto start = std::chrono::system_clock::now();
    constexpr int thread_count = 22;
    Thread threads[thread_count];

    std::filesystem::path p{filename};
    long size = (long)std::filesystem::file_size(p);
    long chunk_size = size / thread_count;

    for (int i = 0; i < thread_count; ++i) {
        threads[i].start = i * chunk_size;
        threads[i].end = (i + 1) * chunk_size - 1;
        if (i == thread_count - 1) {
            threads[i].end = size;
        }
        threads[i].thread = std::thread{process_line, std::ref(threads[i])};
    }

    for (auto & thread : threads) {
        thread.thread.join();
    }

    auto end = std::chrono::system_clock::now();
    typedef std::chrono::duration<float> fsec;
    fsec fs = end - start;
    std::cout << fs.count() << "s\n";
    return 0;
}

int main_ifstream() {
    auto start = std::chrono::system_clock::now();
    char data[0x1000];
    std::ifstream in("measurements.txt");

    while (in)
    {
        in.read(data, 0x1000);
        // do something with data
    }
    in.close();
    auto end = std::chrono::system_clock::now();
    typedef std::chrono::duration<float> fsec;
    fsec fs = end - start;
    std::cout << fs.count() << "s\n";
    return 0;
}

int main_old() {
    auto start = std::chrono::system_clock::now();
    std::ifstream file{"measurements.txt", std::ios_base::in};
    if (!file) std::cerr << "Can't open input file!";
    if (file.is_open()) {
        std::string line;
        int32_t i = 0;
        std::string delimiter = ";";
        std::unordered_map<std::string, locationData> loc_data;
        while (std::getline(file, line)) {
            i++;
            uint64_t pos = line.find(delimiter);
            std::string location = line.substr(0, pos);
            std::string temperature_str = line.substr(pos+1, line.length());
            auto temperature_num = std::stof(temperature_str);
            if (!loc_data.contains(location)) {
                loc_data[location] = locationData{
                    .count = 1,
                    .total = temperature_num,
                    .temp_min = temperature_num,
                    .temp_max = temperature_num,
                };
            } else {
                auto loc_value = loc_data.at(location);
                loc_value.count++;
                loc_value.total += temperature_num;
                if (loc_value.temp_max < temperature_num) {
                    loc_value.temp_max = temperature_num;
                }
                if (loc_value.temp_min > temperature_num) {
                    loc_value.temp_min = temperature_num;
                }
                loc_data[location] = loc_value;
            }
        }
        file.close();
    }
    auto end = std::chrono::system_clock::now();
    typedef std::chrono::duration<float> fsec;
    fsec fs = end - start;
    std::cout << fs.count() << "s\n";
    return 0;
}

int main() {
    return main_mmap();
}