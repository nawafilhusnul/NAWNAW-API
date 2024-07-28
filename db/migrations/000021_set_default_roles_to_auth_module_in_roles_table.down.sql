UPDATE `roles` 
SET `module_id` = NULL 
WHERE `module_id` = (SELECT `id` FROM `modules` WHERE `name` = 'auths' LIMIT 1);