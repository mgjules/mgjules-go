CREATE MIGRATION m1gyedhha4mbbtm5i6qhf6xhfwenia4ignbcosnfxh4q3cdxschn4q
    ONTO m1q2vbhjxpatvt7yc6sojvdqhojio4ojrrrjzfl4l5ogbfj2gio7oa
{
  CREATE TYPE default::User EXTENDING default::Timestamps, default::Avatar {
      CREATE REQUIRED PROPERTY first_name -> std::str;
      CREATE REQUIRED PROPERTY last_name -> std::str;
  };
  CREATE TYPE default::BlogPost EXTENDING default::Timestamps {
      CREATE REQUIRED LINK created_by -> default::User {
          CREATE CONSTRAINT std::exclusive;
      };
      CREATE REQUIRED PROPERTY content -> std::str;
      CREATE REQUIRED PROPERTY cover_image -> std::str;
      CREATE REQUIRED PROPERTY title -> std::str;
      CREATE REQUIRED PROPERTY slug := (std::str_lower(std::str_replace(.title, ' ', '-')));
      CREATE REQUIRED PROPERTY summary -> std::str;
  };
  CREATE TYPE default::BlogTag EXTENDING default::Timestamps {
      CREATE REQUIRED PROPERTY name -> std::str;
      CREATE REQUIRED PROPERTY slug := (std::str_lower(std::str_replace(.name, ' ', '-')));
  };
  ALTER TYPE default::BlogPost {
      CREATE REQUIRED MULTI LINK tags -> default::BlogTag;
  };
};
