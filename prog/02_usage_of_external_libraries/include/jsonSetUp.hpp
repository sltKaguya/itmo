#pragma once

#include <chrono>
#include <ctime>
#include <fstream>
#include <iostream>
#include <string>

#include "json_nlohmann/single_include/nlohmann/json.hpp"
using json = nlohmann::json;

/**
 * @brief Get a string to parse
 * @param jsonString pointer to the string to read from
 * @return Created json
 */
json jsonReader(std::string *jsonString);

/**
 * @brief Print all valutes with their nominals in a file
 */
void printFromJson(json j);