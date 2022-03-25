#include "../include/jsonSetUp.hpp"

json jsonReader(std::string *jsonString) {
    json j = json::parse(*jsonString);

    return j;
}

void printFromJson(json j) {
    std::ofstream file("result.out");

    auto now = std::chrono::system_clock::now();
    std::time_t time = std::chrono::system_clock::to_time_t(now);
    file << "Время запроса: " << std::ctime(&time);

    file << "Последнее обновление данных: " << j["PreviousDate"] << std::endl;
    file << "Следующее обновление данных: " << j["Date"] << std::endl;

    std::cout << j.dump(4) << std::endl;
    //std::string valute = j["Valute"];
    //file << valute;
}

