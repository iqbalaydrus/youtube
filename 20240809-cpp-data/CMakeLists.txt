cmake_minimum_required(VERSION 3.28)
project(20240809_cpp_data)

set(CMAKE_CXX_STANDARD 20)
set(Boost_NO_WARN_NEW_VERSIONS 1)

# Specify the path to the Boost root directory
set(BOOST_ROOT "${CMAKE_SOURCE_DIR}/boost_1_86_0")  # Adjust the version if needed
set(BOOST_INCLUDEDIR "${BOOST_ROOT}/include")
set(BOOST_LIBRARYDIR "${BOOST_ROOT}/lib")

# Find the Boost libraries you need (example: filesystem and system)
find_package(Boost 1.86.0 REQUIRED COMPONENTS iostreams)

# Include Boost headers
include_directories(${Boost_INCLUDE_DIRS})
include_directories("${CMAKE_SOURCE_DIR}/include")

add_executable(20240809_cpp_data main.cpp)

# Link Boost libraries to your executable
target_link_libraries(20240809_cpp_data ${Boost_LIBRARIES})
