INSERT INTO user_roles (
  user_id,
  role_id
) VALUES (
  (SELECT id FROM users WHERE email = 'husnulnawafil27@gmail.com' AND phone = '+6282249907755' LIMIT 1),
  (SELECT id FROM roles WHERE slug = 'super-admin' LIMIT 1)
);