FROM mcr.microsoft.com/dotnet/sdk:8.0 as build-env
WORKDIR /src

COPY . ./

ENV ASPNETCORE_URLS=http://+:55022

RUN dotnet restore
RUN dotnet publish -c release -o /out --no-restore

FROM mcr.microsoft.com/dotnet/aspnet:8.0
WORKDIR /app
COPY --from=build-env /out .

ENTRYPOINT ["dotnet", "NotificationSystem.dll"]
