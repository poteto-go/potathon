#include <iostream>

int main(){
  int a, b; std::cin >> a >> b;
  std::cerr << "Error output:" << a << " " << b << std::endl; // This will output to standard error
  std::cout << a + b << std::endl; // This will output to standard output
  return 0;
}