FROM mcr.microsoft.com/dotnet/sdk:9.0.301-alpine3.22-amd64 AS builder

RUN apk add --no-cache build-base zlib-dev

ENV DOTNET_CLI_TELEMETRY_OPTOUT=1 VERSION=770406a

RUN curl -L https://github.com/SirBogman/Hexagony/tarball/$VERSION | tar xz \
 && cd /SirBogman-Hexagony-$VERSION                                         \
 && sed -Ei 's/\bnet\d+\b/net9/' Hexagony.csproj                            \
 && VERSION= dotnet publish -c Release -r linux-musl-x64                    \
    -p:InvariantGlobalization=true -p:PublishAot=true -p:PublishDir=/       \
    -p:StripSymbols=true

FROM codegolf/lang-base

COPY --from=0 /lib/ld-musl-x86_64.so.1 /lib/
COPY --from=0 /usr/lib/libgcc_s.so.1   \
              /usr/lib/libstdc++.so.6  /usr/lib/
COPY --from=0 /Hexagony                /usr/bin/hexagony

ENTRYPOINT ["hexagony"]
