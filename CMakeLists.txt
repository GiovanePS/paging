cmake_minimum_required(VERSION 3.10)

project(paging)

add_executable(main main.c system/system.c process/process.c memory/memory.c utils/utils.c)
