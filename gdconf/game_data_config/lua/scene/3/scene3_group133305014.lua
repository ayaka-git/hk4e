-- 基础信息
local base_info = {
	group_id = 133305014
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 14001, monster_id = 21010101, pos = { x = -2375.970, y = 217.236, z = 4001.427 }, rot = { x = 0.000, y = 267.661, z = 0.000 }, level = 32, drop_tag = "丘丘人", pose_id = 9017, area_id = 26 },
	{ config_id = 14002, monster_id = 21010101, pos = { x = -2377.601, y = 215.667, z = 4008.373 }, rot = { x = 0.000, y = 12.479, z = 0.000 }, level = 32, drop_tag = "丘丘人", pose_id = 9017, area_id = 26 }
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
	end_suite = 0,
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
		monsters = { 14001, 14002 },
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