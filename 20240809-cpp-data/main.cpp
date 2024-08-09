#include <iostream>
#include <fstream>
#include <map>

struct locationData {
    uint64_t count = 0;
    float_t total = 0;
    float_t temp_min = 99;
    float_t temp_max = -99;
};

void line_process() {

}

int main() {
    auto start = std::chrono::system_clock::now();
    std::ifstream file("measurements-small.txt", std::ios_base::in);
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
            float temperature_num = std::stof(temperature_str);
            if (loc_data.find(location) == loc_data.end()) {
                loc_data[location] = locationData{
                    .count = 1,
                    .total = temperature_num,
                    .temp_min = temperature_num,
                    .temp_max = temperature_num,
                };
            } else {
                loc_data[location].count++;
                loc_data[location].total += temperature_num;
                if (loc_data[location].temp_max < temperature_num) {
                    loc_data[location].temp_max = temperature_num;
                }
                if (loc_data[location].temp_min > temperature_num) {
                    loc_data[location].temp_min = temperature_num;
                }
            }
//            if (3 == i) {
//                break;
//            }
        }
        file.close();
    }
    auto end = std::chrono::system_clock::now();
    typedef std::chrono::duration<float> fsec;
    fsec fs = end - start;
    std::cout << fs.count() << "s\n";
    return 0;
}
