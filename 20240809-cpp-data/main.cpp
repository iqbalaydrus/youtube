#include <chrono>
#include <cmath>
#include <cstdint>
#include <thread>
#include <iostream>
#include <fstream>
#include <unordered_map>
#include <boost/iostreams/device/mapped_file.hpp>
#include <boost/iostreams/device/array.hpp>
#include <boost/iostreams/stream.hpp>

struct locationData {
    uint64_t count = 0;
    float_t total = 0;
    float_t temp_min = 99;
    float_t temp_max = -99;
};

namespace io = boost::iostreams;
constexpr int buf_size = 4096;

struct Thread {
    std::thread thread;
    std::atomic<int> has_data;
    long length{};
    char buffer[buf_size]{};
};

void process_line(Thread &t) {
    std::string line;
    std::string delimiter = ";";
    while (true) {
        t.has_data.wait(0);
        if (t.has_data == 2) {
            break;
        }
        io::basic_array_source<char> input_source(t.buffer, t.length);
        io::stream in(input_source);
        while (std::getline(in, line)) {
            uint64_t pos = line.find(delimiter);
            std::string location = line.substr(0, pos);
            std::string temperature_str = line.substr(pos+1, line.length());
            auto temperature_num = std::stof(temperature_str);
        }
        t.has_data = 0;
        t.has_data.notify_one();
    }
}

int main_mmap() {
    auto start = std::chrono::system_clock::now();
    constexpr int thread_count = 11;
    Thread threads[thread_count];

    for (auto & i : threads) {
        i.thread = std::thread{process_line, std::ref(i)};
    }

    io::stream_buffer<io::mapped_file_source> file("measurements.txt");
    std::istream in(&file);
    auto iter = 0;
    auto fsize = file->size();
    while (in.good())
    {
        auto thread_id = iter%thread_count;
        threads[thread_id].has_data.wait(1);
        in.read(threads[thread_id].buffer, buf_size);
        long i;
        auto gcount = in.gcount();
        auto pos = in.tellg();
        for (i = gcount-1; i >= 0; --i) {
            if (threads[thread_id].buffer[i] == '\n') {
                threads[thread_id].length = i + 1;
                break;
            }
            if (pos >= fsize) {
                threads[thread_id].length = gcount;
                break;
            }
        }
        pos = pos - (gcount - i + 1);
        in.seekg(pos);
        ++iter;
        threads[thread_id].has_data = 1;
        threads[thread_id].has_data.notify_one();
    }
    file.close();
    for (auto & thread : threads) {
        thread.has_data.wait(1);
        thread.has_data = 2;
        thread.has_data.notify_one();
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