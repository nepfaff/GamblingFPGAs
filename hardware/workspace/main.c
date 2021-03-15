#include "src/globals.h"
#include "src/setup.h"
#include "src/request.h"
#include "src/timerLoop.h"

Data data;
DataSrc dataSrc;

int main()
{
	printf("Checking Peripherals..\n");

	// setup peripherals
	if (setupPeripherals()) {
		return 1;
	}

	// setup jtag
	if(openFP()){
		fprintf(stderr, "Unable to access jtag.\n");
		return 1;
	}

	// setup timer
	setupTimerLoop();

	// start main execution
	printf("Running ..\n");

	// run request loop
	requestLoop();

	// close jtag
	closeFP();

	printf("Complete\n");

	return 0;
}
