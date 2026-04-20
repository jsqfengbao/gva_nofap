package source

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model"
	"go.uber.org/zap"
)

// InitAchievements 初始化成就系统基础数据
func InitAchievements() {
	achievements := []model.Achievement{
		// 打卡类成就
		{
			Name:         "初心不改",
			Description:  "完成第一次打卡",
			Category:     1,
			Type:         3,
			Condition:    `{"checkin_count": 1}`,
			Rewards:      10,
			Rarity:       1,
			DisplayOrder: 1,
			IsActive:     true,
		},
		{
			Name:         "坚持一周",
			Description:  "连续打卡7天",
			Category:     1,
			Type:         2,
			Condition:    `{"streak_days": 7}`,
			Rewards:      50,
			Rarity:       1,
			DisplayOrder: 2,
			IsActive:     true,
		},
		{
			Name:         "月度坚持",
			Description:  "连续打卡30天",
			Category:     1,
			Type:         2,
			Condition:    `{"streak_days": 30}`,
			Rewards:      200,
			Rarity:       2,
			DisplayOrder: 3,
			IsActive:     true,
		},
		{
			Name:         "百日重生",
			Description:  "连续打卡100天",
			Category:     1,
			Type:         2,
			Condition:    `{"streak_days": 100}`,
			Rewards:      500,
			Rarity:       3,
			DisplayOrder: 4,
			IsActive:     true,
		},
		{
			Name:         "年度英雄",
			Description:  "连续打卡365天",
			Category:     1,
			Type:         2,
			Condition:    `{"streak_days": 365}`,
			Rewards:      1000,
			Rarity:       4,
			DisplayOrder: 5,
			IsActive:     true,
		},

		// 等级类成就
		{
			Name:         "初级学徒",
			Description:  "达到等级5",
			Category:     2,
			Type:         3,
			Condition:    `{"level": 5}`,
			Rewards:      30,
			Rarity:       1,
			DisplayOrder: 6,
			IsActive:     true,
		},
		{
			Name:         "进阶行者",
			Description:  "达到等级10",
			Category:     2,
			Type:         3,
			Condition:    `{"level": 10}`,
			Rewards:      80,
			Rarity:       1,
			DisplayOrder: 7,
			IsActive:     true,
		},
		{
			Name:         "资深导师",
			Description:  "达到等级25",
			Category:     2,
			Type:         3,
			Condition:    `{"level": 25}`,
			Rewards:      300,
			Rarity:       2,
			DisplayOrder: 8,
			IsActive:     true,
		},
		{
			Name:         "传奇大师",
			Description:  "达到等级50",
			Category:     2,
			Type:         3,
			Condition:    `{"level": 50}`,
			Rewards:      800,
			Rarity:       4,
			DisplayOrder: 9,
			IsActive:     true,
		},

		// 社区类成就
		{
			Name:         "初来乍到",
			Description:  "发表第一篇动态",
			Category:     3,
			Type:         3,
			Condition:    `{"post_count": 1}`,
			Rewards:      15,
			Rarity:       1,
			DisplayOrder: 10,
			IsActive:     true,
		},
		{
			Name:         "热心分享",
			Description:  "累计发表10篇动态",
			Category:     3,
			Type:         1,
			Condition:    `{"post_count": 10}`,
			Rewards:      100,
			Rarity:       1,
			DisplayOrder: 11,
			IsActive:     true,
		},
		{
			Name:         "人气达人",
			Description:  "单篇动态获得100个赞",
			Category:     3,
			Type:         3,
			Condition:    `{"post_likes": 100}`,
			Rewards:      150,
			Rarity:       2,
			DisplayOrder: 12,
			IsActive:     true,
		},
		{
			Name:         "社区之星",
			Description:  "累计获得1000个赞",
			Category:     3,
			Type:         1,
			Condition:    `{"total_likes": 1000}`,
			Rewards:      400,
			Rarity:       3,
			DisplayOrder: 13,
			IsActive:     true,
		},

		// 学习类成就
		{
			Name:         "求知若渴",
			Description:  "完成第一篇学习内容",
			Category:     4,
			Type:         3,
			Condition:    `{"learning_completed": 1}`,
			Rewards:      20,
			Rarity:       1,
			DisplayOrder: 14,
			IsActive:     true,
		},
		{
			Name:         "博学多才",
			Description:  "累计学习时长达到100小时",
			Category:     4,
			Type:         1,
			Condition:    `{"learning_hours": 100}`,
			Rewards:      250,
			Rarity:       2,
			DisplayOrder: 15,
			IsActive:     true,
		},
		{
			Name:         "学者风范",
			Description:  "完成所有类型的学习内容",
			Category:     4,
			Type:         3,
			Condition:    `{"all_types_completed": true}`,
			Rewards:      300,
			Rarity:       3,
			DisplayOrder: 16,
			IsActive:     true,
		},

		// 特殊成就
		{
			Name:         "助人为乐",
			Description:  "帮助他人解决紧急求助10次",
			Category:     5,
			Type:         1,
			Condition:    `{"help_responses": 10}`,
			Rewards:      200,
			Rarity:       2,
			DisplayOrder: 17,
			IsActive:     true,
		},
		{
			Name:         "完美开局",
			Description:  "首次评估风险等级为正常",
			Category:     5,
			Type:         3,
			Condition:    `{"first_assessment_normal": true}`,
			Rewards:      100,
			Rarity:       2,
			DisplayOrder: 18,
			IsActive:     true,
		},
		{
			Name:         "逆转人生",
			Description:  "风险等级从重度降为正常",
			Category:     5,
			Type:         3,
			Condition:    `{"risk_level_improvement": "severe_to_normal"}`,
			Rewards:      600,
			Rarity:       4,
			DisplayOrder: 19,
			IsActive:     true,
		},
		{
			Name:         "钢铁意志",
			Description:  "单月内情绪评分平均达到4分以上",
			Category:     5,
			Type:         3,
			Condition:    `{"monthly_mood_avg": 4.0}`,
			Rewards:      180,
			Rarity:       2,
			DisplayOrder: 20,
			IsActive:     true,
		},
	}

	db := global.GVA_DB
	for _, achievement := range achievements {
		var count int64
		db.Model(&model.Achievement{}).Where("name = ?", achievement.Name).Count(&count)
		if count == 0 {
			err := db.Create(&achievement).Error
			if err != nil {
				global.GVA_LOG.Error("成就数据初始化失败", zap.String("name", achievement.Name), zap.Error(err))
			} else {
				global.GVA_LOG.Info("成就数据初始化成功", zap.String("name", achievement.Name))
			}
		}
	}
}
