package source

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model"
)

// InitLearningContents 初始化学习内容数据
func InitLearningContents() {
	// 检查是否已经存在数据，避免重复插入
	var count int64
	global.GVA_DB.Model(&model.LearningContent{}).Count(&count)
	if count > 0 {
		return
	}

	now := time.Now()
	learningContents := []model.LearningContent{
		// 科普知识类内容
		{
			Title:        "认识色情成瘾的本质",
			Summary:      "深入了解色情成瘾的生理和心理机制，为戒除打下科学基础",
			Content:      `色情成瘾是一种行为成瘾，它会改变大脑的奖赏回路。当我们接触色情内容时，大脑会释放大量多巴胺，形成强化学习机制。长期下去，大脑会对正常的愉悦刺激失去敏感性，需要更强烈的刺激才能获得满足感。这就是为什么戒色如此困难的根本原因。了解这一点，我们就能理解戒色是一个需要时间和耐心的过程，不是简单的意志力问题。`,
			ContentType:  1, // 文章
			Category:     1, // 科普知识
			Difficulty:   1, // 入门
			Duration:     8,
			ThumbnailUrl: "https://images.unsplash.com/photo-1559757148-5c350d0d3c56?w=400&h=300&fit=crop",
			Author:       "心理学专家",
			Tags:         "成瘾机制,多巴胺,神经科学,戒色原理",
			Status:       1,
			PublishAt:    &now,
		},
		{
			Title:        "戒色的生理变化时间表",
			Summary:      "了解戒色过程中身体和大脑的恢复时间表，增强坚持信心",
			Content:      `戒色是一个循序渐进的过程，了解身体的恢复时间表有助于保持耐心和信心。第1-7天：戒断反应期，可能出现焦虑、失眠等症状；第8-30天：适应期，情绪逐渐稳定，注意力开始集中；第31-90天：重建期，大脑开始重新建立健康的奖赏机制；第91-365天：稳定期，新的行为模式逐渐固化；365天以上：完全康复期，大脑功能基本恢复正常。记住，每个人的恢复速度不同，保持耐心是关键。`,
			ContentType:  1,
			Category:     1,
			Difficulty:   2, // 初级
			Duration:     10,
			ThumbnailUrl: "https://images.unsplash.com/photo-1559757175-0eb30cd8c063?w=400&h=300&fit=crop",
			Author:       "康复专家",
			Tags:         "恢复时间表,戒断反应,大脑恢复,康复阶段",
			Status:       1,
			PublishAt:    &now,
		},

		// 康复指导类内容
		{
			Title:        "21天自控力训练计划",
			Summary:      "系统性的自控力训练方法，帮你建立持久的意志力",
			Content:      `自控力是可以通过训练提高的能力。这个21天计划分为三个阶段：第1-7天：建立基础，每天进行5分钟冥想，记录触发因素；第8-14天：强化训练，增加运动量，练习延迟满足；第15-21天：巩固习惯，建立长期的生活方式。具体方法包括：1.每天固定时间冥想 2.冷水澡训练意志力 3.运动释放多巴胺 4.阅读提升认知 5.社交活动转移注意力。坚持21天，你会发现自控力显著提升。`,
			ContentType:  1,
			Category:     2, // 康复指导
			Difficulty:   2,
			Duration:     15,
			ThumbnailUrl: "https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?w=400&h=300&fit=crop",
			Author:       "行为治疗师",
			Tags:         "自控力训练,21天挑战,意志力,行为改变",
			Status:       1,
			PublishAt:    &now,
		},
		{
			Title:        "应对复发的5步法",
			Summary:      "复发是戒色过程中的常见现象，学会正确应对是成功的关键",
			Content:      `复发不等于失败，关键是如何正确应对。5步应对法：第1步：接受现实，不要过度自责；第2步：分析原因，找出触发因素；第3步：调整策略，改进防护措施；第4步：重新开始，制定新的计划；第5步：寻求支持，不要独自承担。记住，大多数成功戒色的人都经历过多次复发，这是正常的学习过程。重要的是从每次复发中学习，不断完善自己的戒色策略。`,
			ContentType:  1,
			Category:     2,
			Difficulty:   3, // 中级
			Duration:     12,
			ThumbnailUrl: "https://images.unsplash.com/photo-1516321318423-f06f85e504b3?w=400&h=300&fit=crop",
			Author:       "康复顾问",
			Tags:         "复发应对,戒色策略,心理调节,重新开始",
			Status:       1,
			PublishAt:    &now,
		},

		// 心理健康类内容
		{
			Title:        "焦虑情绪管理指南",
			Summary:      "学会识别和管理戒色过程中的焦虑情绪，保持心理健康",
			Content:      `戒色初期出现焦虑是正常现象，学会管理焦虑能大大提高成功率。焦虑识别：心跳加快、呼吸急促、坐立不安、思维混乱。管理方法：1.深呼吸练习：4-7-8呼吸法，吸气4秒，憋气7秒，呼气8秒；2.渐进性肌肉放松：从脚趾开始，逐步放松全身肌肉；3.认知重构：质疑负面想法，寻找更客观的观点；4.分散注意力：运动、听音乐、与朋友聊天；5.专业求助：严重时及时寻求心理咨询。`,
			ContentType:  1,
			Category:     3, // 心理健康
			Difficulty:   2,
			Duration:     13,
			ThumbnailUrl: "https://images.unsplash.com/photo-1544367567-0f2fcb009e0b?w=400&h=300&fit=crop",
			Author:       "心理咨询师",
			Tags:         "焦虑管理,深呼吸,肌肉放松,情绪调节",
			Status:       1,
			PublishAt:    &now,
		},
		{
			Title:        "建立积极的自我对话",
			Summary:      "改变内在声音，用积极的自我对话替代负面的自我批评",
			Content:      `内在对话深刻影响我们的情绪和行为。消极的自我对话如"我又失败了"、"我没有意志力"会削弱信心，而积极的自我对话能增强力量。建立积极自我对话的方法：1.识别负面想法：注意内心的批评声音；2.质疑其真实性：这些想法是事实还是假设？3.寻找平衡观点：从朋友的角度看待自己；4.使用肯定语句："我正在学习和成长"、"每一天都是新的开始"；5.练习自我慈悲：像对待好朋友一样对待自己。`,
			ContentType:  1,
			Category:     3,
			Difficulty:   2,
			Duration:     11,
			ThumbnailUrl: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=400&h=300&fit=crop",
			Author:       "正念导师",
			Tags:         "自我对话,积极思维,自我慈悲,心理调节",
			Status:       1,
			PublishAt:    &now,
		},

		// 经验分享类内容
		{
			Title:        "我的180天戒色心路历程",
			Summary:      "一位成功戒色180天的朋友分享真实经历和宝贵经验",
			Content:      `我想分享我180天戒色的真实经历，希望能给正在路上的朋友一些启发。前30天是最困难的，每天都像在打仗。我学会了用运动转移注意力，每当有冲动就去跑步。60天时遇到了第一次严重的冲动波，差点破戒，但我及时联系了戒友，他们的支持帮我度过难关。90天后，我发现自己的专注力明显提升，人际关系也在改善。120天时，戒色已经成为一种生活方式。现在180天了，我感觉重新找回了真实的自己。最重要的经验：1.建立支持网络 2.培养健康爱好 3.保持学习和成长 4.对自己要有耐心。`,
			ContentType:  1,
			Category:     4, // 经验分享
			Difficulty:   1,
			Duration:     18,
			ThumbnailUrl: "https://images.unsplash.com/photo-1522075469751-3a6694fb2f61?w=400&h=300&fit=crop",
			Author:       "戒友小李",
			Tags:         "成功经验,180天,心路历程,戒色技巧",
			Status:       1,
			PublishAt:    &now,
		},
		{
			Title:        "从屡战屡败到成功戒色的转变",
			Summary:      "分享多次失败后终于成功的关键转变点和实用方法",
			Content:      `我戒色失败了无数次，每次都信心满满地开始，然后在几天或几周后崩溃。直到我意识到问题所在：我把戒色当成了一场战争，而不是一次学习。真正的转变发生在我改变心态之后：1.不再追求完美：允许自己有挫折，重点是持续改进；2.建立系统而非依赖意志力：制定详细的计划和应急方案；3.关注进步而非天数：庆祝每一个小的进步；4.寻求专业帮助：阅读相关书籍，参加支持小组；5.改变环境：移除触发因素，创造支持性环境。现在我已经成功戒色8个月，生活发生了翻天覆地的变化。`,
			ContentType:  1,
			Category:     4,
			Difficulty:   3,
			Duration:     16,
			ThumbnailUrl: "https://images.unsplash.com/photo-1531384698654-7f6e477ca221?w=400&h=300&fit=crop",
			Author:       "戒友阿明",
			Tags:         "失败经验,心态转变,系统方法,持续改进",
			Status:       1,
			PublishAt:    &now,
		},

		// 音频内容
		{
			Title:        "10分钟正念冥想引导",
			Summary:      "专业录制的正念冥想音频，帮助你在冲动时保持觉察和平静",
			Content:      `这是一段专业录制的10分钟正念冥想引导音频。通过专注呼吸和身体感觉，帮助你培养觉察力，在面对冲动时保持冷静和理智。适合每天练习，也可以在感到冲动时使用。冥想步骤包括：身体扫描、呼吸觉察、想法观察、慈心练习等。`,
			ContentType:  3, // 音频
			Category:     3,
			Difficulty:   1,
			Duration:     10,
			ThumbnailUrl: "https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop",
			MediaUrl:     "https://example.com/meditation-audio.mp3",
			Author:       "正念导师",
			Tags:         "正念冥想,音频引导,觉察力,冲动管理",
			Status:       1,
			PublishAt:    &now,
		},

		// 视频内容
		{
			Title:        "高效的戒色运动训练",
			Summary:      "15分钟无器械全身训练，释放多巴胺，增强意志力",
			Content:      `运动是戒色的最佳伙伴。这套15分钟的无器械训练专门为戒色朋友设计，包括：热身5分钟、力量训练8分钟、拉伸2分钟。动作简单易学，在家就能完成。规律运动能够：1.释放天然多巴胺 2.提高自信心 3.改善睡眠质量 4.增强体质 5.转移注意力。建议每天或隔天进行，坚持一个月就能看到明显效果。`,
			ContentType:  2, // 视频
			Category:     2,
			Difficulty:   2,
			Duration:     15,
			ThumbnailUrl: "https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?w=400&h=300&fit=crop",
			MediaUrl:     "https://example.com/workout-video.mp4",
			Author:       "健身教练",
			Tags:         "运动训练,无器械,全身训练,多巴胺,意志力",
			Status:       1,
			PublishAt:    &now,
		},
	}

	// 批量插入数据
	if err := global.GVA_DB.Create(&learningContents).Error; err != nil {
		global.GVA_LOG.Error("初始化学习内容失败: " + err.Error())
	} else {
		global.GVA_LOG.Info("学习内容数据初始化成功")
	}
}
