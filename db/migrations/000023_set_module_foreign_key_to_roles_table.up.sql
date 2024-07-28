ALTER TABLE `roles`
ADD CONSTRAINT `fk_roles_module_id`
FOREIGN KEY (`module_id`) REFERENCES `modules` (`id`)
ON DELETE CASCADE
ON UPDATE CASCADE;
