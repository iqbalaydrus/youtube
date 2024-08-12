cmake_minimum_required(VERSION 3.28)
project(20240809_cpp_data)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -stdlib=stdlibc++")

add_executable(20240809_cpp_data main.cpp)
