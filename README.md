# go-checksums

It's a small utility used for getting a checksum for any file.

## Usage

```bash
    go-checksums.exe -f, --file FILE
```

## Demo

```bash
    go run . --file test.txt
    Processing...

    MD5:
    04ff2868ab8b0fd8c5a7075e935b451c

    Sha1:
    319c41f3bc8e4c2df19d60ff04e5bac255e36ffe

    Sha224:
    60428e84740e0eafad716e063ac086cc65593bb88dece289600b5cd1

    Sha256:
    b9bbd2bbe3fd9b9c2121ddfdfda6c068b87a0403672067baa0d4527d836552e7

    Sum384:
    e06e912ca5b3c3546fd4df8142134817868abe4ab4e16c2f77c88a8da055720973975344622e0507bea48b5004cf46ec

    Sha512:
    f1f3a99cde58c1f59b56f707592f761e3f79f529be1e2d4b0c138f413482e83721e37ccf6760bb2661b45963b17b6d557f92fd1d3481c3351c34c7740c701559

    SHA512_224:
    5af1c3f73f18b9849a3493677e39fedd4b98e380863d9ddeb07e5914

    SHA512_256:
    49ecd0363616a4db7c3e36eb4a31e740cc0aad594076e0fac81a3739e1bc601b

    Finished in: 7.5266ms
```

## License

This project is licensed under the [MIT License](https://github.com/go-nerds/go-checksums/blob/main/LICENSE). See the [LICENSE](https://github.com/go-nerds/go-checksums/blob/main/LICENSE) file for details.

