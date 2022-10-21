CREATE MIGRATION m1vc6v6vz7etz7eeq4te2kkddjeqt35zajv6qbf3smuwv5dwh3dxrq
    ONTO m1gyedhha4mbbtm5i6qhf6xhfwenia4ignbcosnfxh4q3cdxschn4q
{
  ALTER TYPE default::BlogPost {
      ALTER PROPERTY slug {
          USING (std::str_lower(std::str_replace(std::str_trim(.title), ' ', '-')));
      };
  };
  ALTER TYPE default::BlogTag {
      ALTER PROPERTY slug {
          USING (std::str_lower(std::str_replace(std::str_trim(.name), ' ', '-')));
      };
  };
};
