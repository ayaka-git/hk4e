-- 基础信息
local base_info = {
	group_id = 133106422
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 422001, monster_id = 25030301, pos = { x = -343.288, y = 317.060, z = 1414.672 }, rot = { x = 0.000, y = 60.035, z = 0.000 }, level = 36, drop_tag = "盗宝团", area_id = 19 },
	{ config_id = 422002, monster_id = 25070101, pos = { x = -351.002, y = 317.215, z = 1416.869 }, rot = { x = 0.000, y = 289.288, z = 0.000 }, level = 36, drop_tag = "盗宝团", pose_id = 3, area_id = 19 },
	{ config_id = 422003, monster_id = 25010701, pos = { x = -352.183, y = 317.337, z = 1418.405 }, rot = { x = 0.000, y = 117.463, z = 0.000 }, level = 36, drop_tag = "盗宝团", pose_id = 9006, area_id = 19 },
	{ config_id = 422004, monster_id = 25020201, pos = { x = -350.205, y = 317.155, z = 1419.230 }, rot = { x = 0.000, y = 219.817, z = 0.000 }, level = 36, drop_tag = "盗宝团", pose_id = 9009, area_id = 19 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 2,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { 422001, 422002, 422003, 422004 },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	},
	{
		-- suite_id = 2,
		-- description = ,
		monsters = { },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================