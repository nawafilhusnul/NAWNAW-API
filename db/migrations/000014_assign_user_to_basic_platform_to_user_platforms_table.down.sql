DELETE FROM `user_platforms` 
WHERE 
`user_id` = (SELECT id FROM users WHERE email = 'husnulnawafil27@gmail.com' AND phone = '+6282249907755' LIMIT 1) AND 
`platform_id` = (SELECT id FROM platforms WHERE slug = 'basic' LIMIT 1);