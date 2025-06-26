package miniprogram

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram"
	"gorm.io/gorm"
)

var EmergencyResources = []miniprogram.EmergencyResource{
	{
		ID:          1,
		Title:       "4-7-8深呼吸练习",
		Type:        1, // 呼吸练习
		Content:     "这是一种简单而有效的呼吸技巧，可以帮助您快速平静下来。\n\n步骤：\n1. 吸气4秒钟\n2. 屏住呼吸7秒钟\n3. 呼气8秒钟\n4. 重复3-4次\n\n这种呼吸方式可以激活副交感神经系统，帮助身体放松。",
		Duration:    180, // 3分钟
		Difficulty:  1,   // 简单
		IsActive:    true,
		UsageCount:  0,
		AvgRating:   0,
		RatingCount: 0,
		Tags:        "呼吸,放松,平静,焦虑",
		Description: "通过控制呼吸节奏帮助快速平静心情的练习",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Title:       "5分钟正念冥想",
		Type:        2, // 冥想指导
		Content:     "正念冥想可以帮助您专注当下，减少负面情绪的影响。\n\n指导：\n1. 找一个安静的地方坐下\n2. 闭上眼睛，专注于呼吸\n3. 当思绪飘散时，温和地把注意力拉回呼吸\n4. 观察身体的感觉，不做判断\n5. 保持5分钟\n\n记住：没有正确或错误的冥想，只要保持观察即可。",
		Duration:    300, // 5分钟
		Difficulty:  2,   // 中等
		IsActive:    true,
		UsageCount:  0,
		AvgRating:   0,
		RatingCount: 0,
		Tags:        "冥想,正念,专注,当下",
		Description: "通过正念练习帮助专注当下，减少负面情绪",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          3,
		Title:       "舒缓自然音乐",
		Type:        3, // 舒缓音乐
		Content:     "聆听自然声音可以帮助放松心情，减少压力和焦虑。这首音乐包含：森林鸟鸣声、溪水流淌声、轻柔的风声。建议戴上耳机，找一个舒适的位置，闭上眼睛聆听。让自然的声音带走您的烦恼。",
		MediaUrl:    "/audio/nature-sounds.mp3",
		Duration:    600, // 10分钟
		Difficulty:  1,   // 简单
		IsActive:    true,
		UsageCount:  0,
		AvgRating:   0,
		RatingCount: 0,
		Tags:        "音乐,自然,放松,治愈",
		Description: "包含森林鸟鸣和溪水声的舒缓自然音乐",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          4,
		Title:       "你比想象中更坚强",
		Type:        4, // 励志文章
		Content:     "亲爱的朋友，如果你正在阅读这篇文章，说明你正在经历一些困难。首先，我想告诉你：你并不孤单。每个人都会遇到挑战和挫折，这是成长的一部分。重要的是要记住：你比你想象的更坚强，成长需要时间，寻求帮助是勇敢的表现，你的努力很有意义。记住：这种感觉会过去的。你有能力度过这个困难时期，并且会变得更加强大。相信自己，你可以的！",
		Duration:    0, // 阅读类型，无固定时长
		Difficulty:  1, // 简单
		IsActive:    true,
		UsageCount:  0,
		AvgRating:   0,
		RatingCount: 0,
		Tags:        "励志,坚强,成长,鼓励",
		Description: "一篇鼓舞人心的文章，提醒你内在的力量",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          5,
		Title:       "快速运动指导",
		Type:        5, // 运动指导
		Content:     "运动可以释放内啡肽，帮助改善心情。这里有一套简单的运动，无需器械：热身1分钟（原地踏步30秒、手臂绕圈30秒），主要运动3分钟（俯卧撑20个、深蹲20个、平板支撑30秒、开合跳20个），放松1分钟（深呼吸30秒、拉伸30秒）。提示：根据自己的体能调整次数，保持匀速呼吸，如有不适立即停止。",
		Duration:    300, // 5分钟
		Difficulty:  2,   // 中等
		IsActive:    true,
		UsageCount:  0,
		AvgRating:   0,
		RatingCount: 0,
		Tags:        "运动,锻炼,内啡肽,活力",
		Description: "5分钟室内运动，帮助释放压力和负面情绪",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          6,
		Title:       "渐进式肌肉放松",
		Type:        2, // 冥想指导
		Content:     "这是一种通过先紧张后放松肌肉来缓解压力的技巧。\n\n**步骤：**\n\n1. **脚部**：绷紧脚趾5秒，然后放松\n2. **小腿**：绷紧小腿肌肉5秒，然后放松\n3. **大腿**：绷紧大腿肌肉5秒，然后放松\n4. **臀部**：绷紧臀部肌肉5秒，然后放松\n5. **腹部**：绷紧腹部肌肉5秒，然后放松\n6. **胸部**：深吸气绷紧胸部5秒，然后放松\n7. **手臂**：握拳绷紧手臂5秒，然后放松\n8. **肩膀**：耸肩绷紧5秒，然后放松\n9. **面部**：皱眉绷紧面部5秒，然后放松\n\n每次放松后，感受肌肉的松弛感。整个过程约10分钟。",
		Duration:    600, // 10分钟
		Difficulty:  2,   // 中等
		IsActive:    true,
		UsageCount:  0,
		AvgRating:   0,
		RatingCount: 0,
		Tags:        "放松,肌肉,压力,紧张",
		Description: "通过肌肉紧张和放松循环来缓解身体压力",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          7,
		Title:       "情绪稳定呼吸法",
		Type:        1, // 呼吸练习
		Content:     "当情绪激动时，这种呼吸法可以帮助快速稳定情绪。\n\n**方法：**\n\n1. **感受位置**：把一只手放在胸部，一只手放在腹部\n2. **深吸气**：用鼻子慢慢吸气4秒，让腹部的手上升，胸部的手保持相对静止\n3. **暂停**：屏住呼吸2秒\n4. **慢呼气**：用嘴慢慢呼气6秒，腹部的手下降\n5. **重复**：继续这个循环5-10次\n\n**重点：**\n- 专注于腹部呼吸，不是胸部\n- 呼气时间比吸气时间长\n- 保持缓慢而平稳的节奏",
		Duration:    120, // 2分钟
		Difficulty:  1,   // 简单
		IsActive:    true,
		UsageCount:  0,
		AvgRating:   0,
		RatingCount: 0,
		Tags:        "呼吸,情绪,稳定,腹式呼吸",
		Description: "专门针对情绪激动时的腹式呼吸稳定法",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          8,
		Title:       "每个选择都是新的开始",
		Type:        4, // 励志文章
		Content:     "**致正在努力的你：**\n\n生活中最美好的事情之一，就是每一刻都给了我们重新选择的机会。\n\n🔄 **重新开始的力量**\n无论过去发生了什么，现在这一刻，你都可以做出不同的选择。过去不等于未来。\n\n🌅 **每天都是新的一天**\n昨天的挫折不会定义今天的可能性。每个日出都带来新的希望。\n\n🎯 **专注于下一步**\n不要被过去的错误压垮，专注于你现在可以采取的积极行动。\n\n💖 **自我宽恕的重要性**\n学会原谅自己，这是治愈和成长的第一步。你值得第二次机会。\n\n🌟 **你的故事还在续写**\n你的人生故事远未结束。今天的选择将开启新的篇章。\n\n记住：改变从决定开始，而你有能力在任何时候做出这个决定。",
		Duration:    0, // 阅读类型
		Difficulty:  1, // 简单
		IsActive:    true,
		UsageCount:  0,
		AvgRating:   0,
		RatingCount: 0,
		Tags:        "选择,新开始,希望,宽恕",
		Description: "关于重新开始和自我宽恕的励志文章",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

// InitializeEmergencyResources 初始化紧急求助资源
func InitializeEmergencyResources(db *gorm.DB) error {
	// 检查是否已经初始化过
	var count int64
	db.Model(&miniprogram.EmergencyResource{}).Count(&count)
	if count > 0 {
		return nil // 已经初始化过，跳过
	}

	// 批量插入资源数据
	return db.Create(&EmergencyResources).Error
}
