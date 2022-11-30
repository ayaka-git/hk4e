-- 基础信息
local base_info = {
	group_id = 133002212
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 1583, gadget_id = 70210101, pos = { x = 1954.237, y = 219.321, z = -958.042 }, rot = { x = 0.000, y = 323.354, z = 0.000 }, level = 6, drop_tag = "搜刮点解谜通用蒙德", area_id = 3 },
	{ config_id = 1584, gadget_id = 70210101, pos = { x = 1917.644, y = 224.611, z = -939.219 }, rot = { x = 0.000, y = 323.354, z = 0.000 }, level = 6, drop_tag = "搜刮点解谜通用蒙德", area_id = 3 },
	{ config_id = 1585, gadget_id = 70210101, pos = { x = 1916.543, y = 224.080, z = -887.045 }, rot = { x = 0.000, y = 323.354, z = 0.000 }, level = 6, drop_tag = "搜刮点解谜通用蒙德", area_id = 3 },
	{ config_id = 1586, gadget_id = 70210101, pos = { x = 1931.042, y = 219.799, z = -877.258 }, rot = { x = 0.000, y = 323.354, z = 0.000 }, level = 6, drop_tag = "搜刮点解谜通用蒙德", area_id = 3 },
	{ config_id = 1587, gadget_id = 70210101, pos = { x = 1940.772, y = 219.921, z = -874.684 }, rot = { x = 0.000, y = 323.354, z = 0.000 }, level = 6, drop_tag = "搜刮点解谜通用蒙德", area_id = 3 },
	{ config_id = 1588, gadget_id = 70210101, pos = { x = 1922.150, y = 206.168, z = -769.857 }, rot = { x = 0.000, y = 323.354, z = 0.000 }, level = 6, drop_tag = "搜刮点解谜通用蒙德", area_id = 3 },
	{ config_id = 1589, gadget_id = 70210101, pos = { x = 1924.551, y = 210.331, z = -706.622 }, rot = { x = 0.000, y = 323.354, z = 0.000 }, level = 6, drop_tag = "搜刮点解谜通用蒙德", area_id = 3 },
	{ config_id = 1600, gadget_id = 70210101, pos = { x = 1908.641, y = 227.939, z = -870.083 }, rot = { x = 0.000, y = 323.354, z = 0.000 }, level = 6, drop_tag = "搜刮点解谜通用蒙德", area_id = 3 }
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
	rand_suite = true
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
		monsters = { },
		gadgets = { 1583, 1584, 1585, 1586, 1587, 1588, 1589, 1600 },
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