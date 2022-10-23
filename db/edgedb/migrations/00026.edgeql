CREATE MIGRATION m1t32qk3cy4jel7wnd5yai5de55yxpe26dtxiw36giet7kqgq5stxq
    ONTO m1vc6v6vz7etz7eeq4te2kkddjeqt35zajv6qbf3smuwv5dwh3dxrq
{
  ALTER TYPE default::BlogPost {
      ALTER PROPERTY cover_image {
          SET default := 'https://mgjules.dev/img/blog/modern-code-screen.webp';
      };
  };
};
