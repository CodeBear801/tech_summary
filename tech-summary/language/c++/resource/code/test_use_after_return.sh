clang -o use_after_return -g -fsanitize=address use_after_return.cc
ASAN_OPTIONS=detect_stack_use_after_return=1 ./use_after_return