CREATE MIGRATION m1qasyfc5zyuqgx5y4gwe4qclq5ruyttugagjiuymgajawbbwxovta
    ONTO m1h3g6wnbgngkh7lcnfchypf4pwwcn2t4fu3ozzq64ggyj7vk35gsq
{
  ALTER TYPE default::BlogPost {
      ALTER PROPERTY slug {
          USING (std::re_replace('[^a-z]+', '_', std::str_lower(.title), flags := 'g'));
      };
  };
  ALTER TYPE default::BlogTag {
      ALTER PROPERTY slug {
          USING (std::re_replace('[^a-z]+', '_', std::str_lower(.name), flags := 'g'));
      };
  };
};
