FROM mcr.microsoft.com/dotnet/sdk:8.0 as build-env
WORKDIR /src

COPY . ./

RUN dotnet restore
RUN dotnet build -o /out

FROM mcr.microsoft.com/dotnet/aspnet:8.0
WORKDIR /app
COPY --from=build-env /out .

CMD ["./Robot"]
