--> Return spots which have a domain with a count greater than 1.
SELECT * FROM MY_TABLE WHERE rating  >  1;

--> Change the website field, so it only contains the domain.

select substring_index(substring_index(substring_index(website, "/", 3), "//", -1), ".", 3) from MY_TABLE

--> Count how many spots have the same domain.


SELECT substring_index(substring_index(substring_index(website, "/", 3), "//", -1), ".", 3), COUNT(*)
FROM MY_TABLE
GROUP BY substring_index(substring_index(substring_index(website, "/", 3), "//", -1), ".", 3)

--> Return 3 columns: spot name, domain, and count number for domain.

select name, substring_index(substring_index(substring_index(website, "/", 3), "//", -1), ".", 3), COUNT(*)
from MY_TABLE
GROUP BY name, substring_index(substring_index(substring_index(website, "/", 3), "//", -1), ".", 3)
