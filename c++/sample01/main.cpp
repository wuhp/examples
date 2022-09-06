#include "inc.h"
#include "cal.h"
#include "combine.h"
#include <iostream>
using namespace std;

int main(int argc, char** argv) {
  cout << "Input: 1, 2" << endl;
  cout << inc(2) << endl;
  cout << add(1, 2) << endl;
  cout << merge(1, 2) << endl;
  return 0;
}
