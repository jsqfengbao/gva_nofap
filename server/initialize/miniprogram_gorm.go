package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram"
	sourceInit "github.com/flipped-aurora/gin-vue-admin/server/source/miniprogram"
	"go.uber.org/zap"
)

// MiniProgramTables 小程序相关数据表初始化
func MiniProgramTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		// 用户相关表
		&miniprogram.WxUser{},
		&miniprogram.UserSettings{},
		&miniprogram.DataExport{},

		// 戒色记录相关表
		&miniprogram.AbstinenceRecord{},
		&miniprogram.DailyCheckin{},
		&miniprogram.AssessmentResult{},

		// 社区相关表
		&miniprogram.CommunityPost{},
		&miniprogram.CommunityComment{},
		&miniprogram.CommunityLike{},

		// 紧急求助相关表
		&miniprogram.EmergencyHelp{},
		&miniprogram.EmergencyResponse{},
		&miniprogram.EmergencyVolunteer{},
		&miniprogram.EmergencyResource{},

		// 学习内容相关表
		&miniprogram.LearningContent{},
		&miniprogram.UserLearningRecord{},

		// 成就系统相关表
		&miniprogram.Achievement{},
		&miniprogram.UserAchievement{},
	)

	if err != nil {
		global.GVA_LOG.Error("小程序数据表初始化失败!", zap.Error(err))
		panic(err)
	}

	global.GVA_LOG.Info("小程序数据表初始化成功!")

	// 创建必要的索引
	createIndexes()

	// 初始化基础数据
	initBaseData()
}

// initBaseData 初始化基础数据
func initBaseData() {
	// 初始化成就系统数据
	sourceInit.InitAchievements()

	// 初始化紧急求助资源数据
	if err := sourceInit.InitializeEmergencyResources(global.GVA_DB); err != nil {
		global.GVA_LOG.Error("紧急求助资源初始化失败!", zap.Error(err))
	} else {
		global.GVA_LOG.Info("紧急求助资源初始化完成!")
	}

	// 初始化学习内容数据
	sourceInit.InitLearningContents()

	global.GVA_LOG.Info("小程序基础数据初始化完成!")
}

// createIndexes 创建数据库索引优化查询性能
func createIndexes() {
	db := global.GVA_DB

	// 用户表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_wx_users_openid ON nofap_wx_users(openid)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_wx_users_status ON nofap_wx_users(status)")

	// 用户设置表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_user_settings_user_id ON nofap_user_settings(user_id)")

	// 数据导出表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_data_exports_user_status ON nofap_data_exports(user_id, status)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_data_exports_created_at ON nofap_data_exports(created_at)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_data_exports_expires_at ON nofap_data_exports(expires_at)")

	// 戒色记录表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_abstinence_records_user_status ON abstinence_records(user_id, status)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_abstinence_records_level ON abstinence_records(level)")

	// 每日打卡表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_daily_checkins_user_date ON daily_checkins(user_id, checkin_date)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_daily_checkins_date ON daily_checkins(checkin_date)")

	// 评估结果表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_assessment_results_user_date ON assessment_results(user_id, test_date)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_assessment_results_risk_level ON assessment_results(risk_level)")

	// 社区动态表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_community_posts_category_status ON community_posts(category, status)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_community_posts_created_at ON community_posts(created_at)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_community_posts_like_count ON community_posts(like_count)")

	// 社区评论表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_community_comments_post_status ON community_comments(post_id, status)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_community_comments_parent_id ON community_comments(parent_id)")

	// 紧急求助表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_emergency_help_type_status ON nofap_emergency_help(type, status)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_emergency_help_created_at ON nofap_emergency_help(created_at)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_emergency_volunteer_online_status ON nofap_emergency_volunteer(is_online, status)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_emergency_resource_type_active ON nofap_emergency_resource(type, is_active)")

	// 学习内容表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_learning_contents_category_status ON nofap_learning_contents(category, status)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_learning_contents_type ON nofap_learning_contents(content_type)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_learning_contents_view_count ON nofap_learning_contents(view_count)")

	// 用户学习记录表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_user_learning_records_user_content ON nofap_user_learning_records(user_id, content_id)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_user_learning_records_completed ON nofap_user_learning_records(is_completed)")

	// 成就系统表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_achievements_category_active ON achievements(category, is_active)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_achievements_rarity ON achievements(rarity)")

	// 用户成就表索引
	db.Exec("CREATE INDEX IF NOT EXISTS idx_user_achievements_user_unlocked ON user_achievements(user_id, unlocked_at)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_user_achievements_notified ON user_achievements(is_notified)")

	global.GVA_LOG.Info("数据库索引创建完成!")
}
