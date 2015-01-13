if /linux/ =~ RUBY_PLATFORM
  open("Makefile", "wb") do |f|
    f.write <<-EOF
CXX = g++
CXXFLAGS = -O3 -g -Wall

inotify-wrapper: inotify-wrapper.o
	$(CXX) $(CXXFLAGS) $< -o $@

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -o $@

BUILD_DIR=../../build
BINS=zeus-darwin-amd64 zeus-linux-386 zeus-linux-amd64

$(BINS:%=$(BUILD_DIR)/%) :
	@go version || { echo "You need Golang installed to build zeus from source"; exit 1; }
	@echo $@
	make -C ../../.. linux

install: $(BINS:%=$(BUILD_DIR)/%)
    EOF
  end
else
  open("Makefile", "wb") do |f|
    f.write <<-EOF
install:
	# do nothing
    EOF
  end
end
