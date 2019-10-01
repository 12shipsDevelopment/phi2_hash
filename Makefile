CC=gcc
TARGET=libphi2_hash.so

$(TARGET): 
	$(CC) phi2.c lyra2/Lyra2.c sph/jh.c sph/skein.c sph/aes_helper.c sph/cubehash.c sph/streebog.c sph/echo.c lyra2/Sponge.c -fPIC -o $@ -shared
	cp $@ /usr/lib

.PHONY: clean
clean:
	rm -Rf $(OBJ) *.so 

