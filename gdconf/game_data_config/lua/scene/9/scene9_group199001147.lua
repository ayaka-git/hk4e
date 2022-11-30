-- 基础信息
local base_info = {
	group_id = 199001147
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 147004, monster_id = 23010601, pos = { x = 739.922, y = 201.688, z = 202.406 }, rot = { x = 0.000, y = 206.029, z = 0.000 }, level = 20, drop_tag = "先遣队", disableWander = true, pose_id = 9002, area_id = 402 },
	{ config_id = 147005, monster_id = 21010101, pos = { x = 748.378, y = 193.333, z = 221.051 }, rot = { x = 0.000, y = 218.147, z = 0.000 }, level = 20, drop_tag = "丘丘人", disableWander = true, pose_id = 9016, area_id = 402 },
	{ config_id = 147006, monster_id = 23010601, pos = { x = 750.707, y = 201.678, z = 210.233 }, rot = { x = 0.000, y = 7.353, z = 0.000 }, level = 20, drop_tag = "先遣队", disableWander = true, pose_id = 9002, area_id = 402 },
	{ config_id = 147007, monster_id = 23010601, pos = { x = 730.329, y = 185.333, z = 200.223 }, rot = { x = 0.000, y = 232.327, z = 0.000 }, level = 20, drop_tag = "先遣队", disableWander = true, area_id = 402 }
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

-- 废弃数据
garbages = {
	monsters = {
		{ config_id = 147001, monster_id = 21010101, pos = { x = 732.608, y = 179.435, z = 202.560 }, rot = { x = 0.000, y = 103.562, z = 0.000 }, level = 20, drop_tag = "丘丘人", disableWander = true, pose_id = 9016, area_id = 402 }
	}
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
		monsters = { 147004, 147005, 147006, 147007 },
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