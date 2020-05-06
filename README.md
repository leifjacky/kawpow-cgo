### usage

```bash
git clone https://github.com/leifjacky/kawpow-cgo
cd kawpow-cgo/src
mkdir build && cd build
cmake .. && make
cd ../../
ln -fs src/build/lib lib

go test -v ./...
```


