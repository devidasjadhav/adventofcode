# Define Go command and flags
GO = go
GOFLAGS = -ldflags="-s -w"

# Define the target executable
TARGET = aoc

# Default target: build the executable
all: $(TARGET)

# Rule to build the target executable
$(TARGET): clean
	$(GO) build $(GOFLAGS) -o $(TARGET) src/$(dayNumber)/main.go
#	$(GO) build -o $(TARGET) src/1/main.go

# Clean target: remove the target executable
clean:
	rm -f $(TARGET)

# Test target: run Go tests for the project
getinput:
	./getinput.py $(dayNumber) > src/$(dayNumber)/input

temptest: $(TARGET)
	./$(TARGET) < src/$(dayNumber)/input.txt

test: $(TARGET)
	./$(TARGET) < src/$(dayNumber)/input
