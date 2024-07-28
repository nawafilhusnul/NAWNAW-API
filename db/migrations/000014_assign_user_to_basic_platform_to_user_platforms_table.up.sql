INSERT INTO `user_platforms` (`user_id`, `platform_id`) VALUES (
  (SELECT id FROM users WHERE email = 'husnulnawafil27@gmail.com' AND phone = '+6282249907755' LIMIT 1),
  (SELECT id FROM platforms WHERE slug = 'basic' LIMIT 1)
);