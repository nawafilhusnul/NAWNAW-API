UPDATE `roles` 
SET `module_id` = (SELECT `id` FROM `modules` WHERE `name` = 'auths' LIMIT 1)
WHERE `module_id` IS NULL;