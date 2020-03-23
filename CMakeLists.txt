cmake_minimum_required(VERSION 3.5)

project(ttcp LANGUAGES CXX)

set(APP_EXE ttcp)

set(TTCP_HEADERS common.h inet_address.h socket.h)

set(TTCP_SRCS inet_address.cc)

add_library(ttpclib STATIC  ${TTCP_SRCS} ${TTCP_HEADERS})

add_executable(${APP_EXE} ttcp.cc)

target_link_libraries(${APP_EXE} ttpclib)