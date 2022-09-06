#include "inc.h"
#include "lib/cal/cal.h"
#include <iostream>

using namespace std;

int main(int argc, char** argv) {
  int x = 1;
  cout << "Input: " << x << endl;
  cout << "Increased to " << inc(x) << endl;
  cout << "Add 1 and 2 equals " << add(1, 2) << endl;
  return 0;
}
