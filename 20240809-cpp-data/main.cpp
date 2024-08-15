#include <chrono>
#include <cmath>
#include <cstdint>
#include <iostream>
#include <fstream>
#include <unordered_map>
#include <boost/iostreams/device/mapped_file.hpp>
#include <boost/iostreams/stream.hpp>

struct locationData {
    uint64_t count = 0;
    float_t total = 0;
    float_t temp_min = 99;
    float_t temp_max = -99;
};

namespace io = boost::iostreams;

int main_mmap() {
    using namespace boost::iostreams;
    auto start = std::chrono::system_clock::now();
    char data[0x1000];

    io::stream_buffer<io::mapped_file_source> file("measurements.txt");
    std::istream in(&file);
    while (in)
    {
        in.read(data, 0x1000);
        // do something with data
    }
    file.close();

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