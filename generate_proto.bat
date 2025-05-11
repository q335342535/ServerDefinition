@echo off
chcp 65001>nul
echo 开始生成Protocol Buffers文件...

REM 生成route_server相关proto文件

protoc --go_out=. source/common/message.define.proto
if %errorlevel% neq 0 (
    echo 生成common/message.define.proto失败！
    pause
    exit /b %errorlevel%
)

protoc --go_out=. source/route_server/errors.proto
if %errorlevel% neq 0 (
    echo 生成route_server/errors.proto失败！
    pause
    exit /b %errorlevel%
)

protoc --go_out=. source/route_server/protobuf.proto
if %errorlevel% neq 0 (
    echo 生成route_server/protobuf.proto失败！
    pause
    exit /b %errorlevel%
)

REM 生成user_server相关proto文件

protoc --go_out=. source/user_server/errors.proto
if %errorlevel% neq 0 (
    echo 生成user_server/errors.proto失败！
    pause
    exit /b %errorlevel%
)

protoc --go_out=. source/user_server/protobuf.proto
if %errorlevel% neq 0 (
    echo 生成user_server/protobuf.proto失败！
    pause
    exit /b %errorlevel%
)

REM 生成entry_server相关proto文件

protoc --go_out=. source/entry_server/errors.proto
if %errorlevel% neq 0 (
    echo 生成entry_server/errors.proto失败！
    pause
    exit /b %errorlevel%
)

protoc --go_out=. source/entry_server/protobuf.proto
if %errorlevel% neq 0 (
    echo 生成entry_server/protobuf.proto失败！
    pause
    exit /b %errorlevel%
)

REM 生成prop_server相关proto文件

protoc --go_out=. source/prop_server/errors.proto
if %errorlevel% neq 0 (
    echo 生成prop_server/errors.proto失败！
    pause
    exit /b %errorlevel%
)   

protoc --go_out=. source/prop_server/protobuf.proto
if %errorlevel% neq 0 (
    echo 生成prop_server/protobuf.proto失败！
    pause
    exit /b %errorlevel%
)
echo 所有Proto文件生成成功！
pause 