package myorm

// +---------------------------------------------------------------------
// | Description: gorm封装,包含日志记录功能
// +---------------------------------------------------------------------
// | Copyright (c) 2004-2020 护卫神(http://hws.com) All rights reserved.
// +---------------------------------------------------------------------
// | Author: Wjinlei <1976883731@qq.com>
// +---------------------------------------------------------------------
//
//                  ___====-_  _-====___
//             _--^^^#####/      \#####^^^--_
//          _-^##########/ (    ) \##########^-_
//         -############/  |\^^/|  \############-
//       _/############/   (@::@)   \############\_
//     /#############((     \  /     ))#############\
//     -###############\    (oo)    /###############-
//    -#################\  / VV \  /#################-
//   -###################\/      \/###################-
// _#/|##########/\######(   /\   )######/\##########|\#_
// |/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
// '  |/  V  V      V  \#\| |  | |/#/  V      V  V  \|  '
//    '   '  '      '   / | |  | | \   '      '  '   '
//                     (  | |  | |  )
//                    __\ | |  | | /__
//                   (vvv(VVV)(VVV)vvv)
//
//                  神龙护体
//                代码无bug!

import (
	"database/sql"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Option 选项
type Option struct {
	DataSource string          // 数据源
	LogMode    bool            // 是否启用日志
	LogPath    string          // 日志文件路径
	LogLevel   logger.LogLevel // 如果不传,默认为Silent(不记录日志)
}

// DB gorm.DB实例
type DB struct {
	Instance *gorm.DB
	Conn     *sql.DB
}

var (
	err error

	sqliteRotator *rotatelogs.RotateLogs // sqlite 全局日志切割器
	mysqlRotator  *rotatelogs.RotateLogs // mysql 全局日志切割器
)
