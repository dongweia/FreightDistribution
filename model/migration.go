package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&User{},
		&Chatlist{},
		&Chatmess{},
		&Cos{},
		&Friend{},
		&Commodity{},
		&Collect{},
		&WatchHistory{},
		&Order{},
		&Transportlog{},
		&Admin{},
		&Opinion{},)}
