CREATE MIGRATION m12p3oan7mipseu55gdb7snlpky3belg2mvhuavoqeqz7ayutt5kva
    ONTO m1njgyyvp2gx5mqjrkxpcbbq6xpymu7xqliy2vel6acokyg4qovhea
{
  CREATE ABSTRACT LINK default::link_with_sort {
      CREATE PROPERTY sort -> std::int32;
      CREATE CONSTRAINT std::expression ON ((__subject__@sort >= 0));
  };
  ALTER TYPE default::CVExperience {
      ALTER LINK technologies {
          EXTENDING default::link_with_sort LAST;
          ALTER CONSTRAINT std::expression ON ((__subject__@sort >= 0)) {
              DROP OWNED;
          };
          ALTER PROPERTY sort {
              DROP OWNED;
              RESET TYPE;
          };
      };
  };
  CREATE TYPE default::CVProjects EXTENDING default::Timestamps {
      CREATE MULTI LINK technologies EXTENDING default::link_with_sort -> default::CVTechnology;
      CREATE REQUIRED PROPERTY description -> std::str;
      CREATE REQUIRED PROPERTY link -> std::str;
      CREATE REQUIRED PROPERTY name -> std::str {
          CREATE CONSTRAINT std::exclusive;
      };
  };
};
