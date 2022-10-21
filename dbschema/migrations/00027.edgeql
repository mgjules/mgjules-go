CREATE MIGRATION m1h3g6wnbgngkh7lcnfchypf4pwwcn2t4fu3ozzq64ggyj7vk35gsq
    ONTO m1t32qk3cy4jel7wnd5yai5de55yxpe26dtxiw36giet7kqgq5stxq
{
  ALTER TYPE default::BlogPost {
      CREATE INDEX ON (.slug);
  };
  ALTER TYPE default::BlogTag {
      CREATE INDEX ON (.slug);
  };
};
