#include <chrono>
#include <cstdint>
#include <thread>
#include <iostream>
#include <fstream>
#include <unordered_map>
#include <filesystem>
#include <boost/iostreams/device/mapped_file.hpp>
#include <boost/iostreams/stream.hpp>
#include <boost/utility/string_ref.hpp>

struct locationData {
    uint64_t count = 0;
    double total = 0;
    double temp_min = 99;
    double temp_max = -99;
};

namespace io = boost::iostreams;
const std::string filename = "measurements.txt";
constexpr int thread_count = 22;

struct Thread {
    std::thread thread;
    uint8_t num{};
    long start{};
    long end{};
    std::unordered_map<std::string, locationData, std::hash<std::string>> map{};
};


double atof_fast(const char* str) {
    bool neg = false;
    if (*str == '-') {
        neg = true;
        ++str;
    } else if (*str == '+')
        ++str;

    double value = 0;
    for (; *str != '.'; ++str) {
        if (!*str)
            return neg ? -value : value;
        value *= 10;
        value += *str - '0';
    }

    double decimal = 0, weight = 1;
    for (; *++str; weight *= 10) {
        decimal *= 10;
        decimal += *str - '0';
    }
    decimal /= weight;
    return neg ? -(value + decimal) : (value + decimal);
}

void process_line(Thread &t) {
    std::string line;
    line.reserve(128);
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
    uint64_t size = 0;
    while (std::getline(in, line)) {
        ++i;
        size += line.length()+1;
        uint64_t pos = line.rfind(';', line.length()-4);
        auto temperature_str = boost::string_ref{line.data()+pos+1, line.length()-pos-1};
        auto temperature_num = atof_fast(temperature_str.begin());
        auto location_str = std::string{line.data(), pos};
        try {
            auto loc_value = t.map.at(location_str);
            loc_value.count++;
            loc_value.total += temperature_num;
            if (loc_value.temp_max < temperature_num) {
                loc_value.temp_max = temperature_num;
            }
            if (loc_value.temp_min > temperature_num) {
                loc_value.temp_min = temperature_num;
            }
            t.map[location_str] = loc_value;
        } catch (std::out_of_range &e) {
            t.map[location_str] = locationData{
                    .count = 1,
                    .total = temperature_num,
                    .temp_min = temperature_num,
                    .temp_max = temperature_num,
            };
        }
        if (i < 15 && t.num == 0) {
//            std::cout << location_str << " " << temperature_str << std::endl;
        }
        if (t.start + size >= t.end) {
            break;
        }
    }
    file.close();
}

int main_mmap() {
    auto start = std::chrono::system_clock::now();
    Thread threads[thread_count];

    std::filesystem::path p{filename};
    long size = (long)std::filesystem::file_size(p);
    long chunk_size = size / thread_count;

    for (int i = 0; i < thread_count; ++i) {
        threads[i].num = i;
        threads[i].start = i * chunk_size;
        threads[i].end = (i + 1) * chunk_size - 1;
        if (i == thread_count - 1) {
            threads[i].end = size;
        }
        threads[i].thread = std::thread{process_line, std::ref(threads[i])};
        threads[i].map.reserve(512);
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
    std::ifstream in(filename);

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
    std::ifstream file{filename, std::ios_base::in};
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