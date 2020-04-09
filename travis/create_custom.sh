DIR=`pwd`
cd $DIR/build/cmd/bood && 
bood && 
rm $GOPATH/bin/bood &&
mv $DIR/build/cmd/bood/out/bin/bood $GOPATH/bin/ &&
# Create archive package
cd ../../../examples/archive && 
# Archive files
bood 1.txt 2.txt 3.txt &&
# Create binary package
cd ../binary/ &&
# Make out
bood
