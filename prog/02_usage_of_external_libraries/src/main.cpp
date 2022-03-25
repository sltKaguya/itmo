#include "../include/curlSetUp.hpp"
#include "../include/jsonSetUp.hpp"

int main() {
    char url[] = "https://www.cbr-xml-daily.ru/daily_json.js";
    std::string buffer;

    curlSetUp(url, &buffer);

    json j = jsonReader(&buffer);
    printFromJson(j);

    return 0;
}
