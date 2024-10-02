MAIN_BUILD_DIR="./build"
FRONTEND_DIR="./frontend"
SERVER_DIR="./backend/server"
FRONTEND_BUILD_DIR="$FRONTEND_DIR/dist"

# Build frontend with Bun
echo "Building frontend..."
cd $FRONTEND_DIR && bun i && bun run build
cd ..

# Create build directory
echo "Creating build directory..."
mkdir -p $MAIN_BUILD_DIR

# Copy frontend build files
echo "Copying frontend build files..."
cp -r $FRONTEND_BUILD_DIR $MAIN_BUILD_DIR/frontend

# Build Go project
echo "Building Go project..."
go build -o $MAIN_BUILD_DIR/ascii-converter $SERVER_DIR/main.go

echo "Build complete."
