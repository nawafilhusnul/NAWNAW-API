DELETE FROM user_roles WHERE 
user_id = (
  SELECT id FROM users WHERE email = 'husnulnawafil27@gmail.com' AND phone = '+6282249907755'
)
AND role_id = (
  SELECT id FROM roles WHERE slug = 'super-admin'
);