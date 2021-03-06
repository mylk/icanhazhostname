package main

import (
    "fmt"
    "math/rand"
)

var adjectives []string = []string{
    "admiring",
    "adoring",
    "affectionate",
    "agitated",
    "amazing",
    "angry",
    "awesome",
    "beautiful",
    "blissful",
    "bold",
    "boring",
    "brave",
    "busy",
    "charming",
    "clever",
    "cool",
    "compassionate",
    "competent",
    "condescending",
    "confident",
    "cranky",
    "crazy",
    "dazzling",
    "determined",
    "distracted",
    "dreamy",
    "eager",
    "ecstatic",
    "elastic",
    "elated",
    "elegant",
    "eloquent",
    "epic",
    "exciting",
    "fervent",
    "festive",
    "flamboyant",
    "focused",
    "friendly",
    "frosty",
    "funny",
    "gallant",
    "gifted",
    "goofy",
    "gracious",
    "great",
    "happy",
    "hardcore",
    "heuristic",
    "hopeful",
    "hungry",
    "infallible",
    "inspiring",
    "interesting",
    "intelligent",
    "jolly",
    "jovial",
    "keen",
    "kind",
    "laughing",
    "loving",
    "lucid",
    "magical",
    "mystifying",
    "modest",
    "musing",
    "naughty",
    "nervous",
    "nice",
    "nifty",
    "nostalgic",
    "objective",
    "optimistic",
    "peaceful",
    "pedantic",
    "pensive",
    "practical",
    "priceless",
    "quirky",
    "quizzical",
    "recursing",
    "relaxed",
    "reverent",
    "romantic",
    "sad",
    "serene",
    "sharp",
    "silly",
    "sleepy",
    "stoic",
    "strange",
    "stupefied",
    "suspicious",
    "sweet",
    "tender",
    "thirsty",
    "trusting",
    "unruffled",
    "upbeat",
    "vibrant",
    "vigilant",
    "vigorous",
    "wizardly",
    "wonderful",
    "xenodochial",
    "youthful",
    "zealous",
    "zen",
}

type Person struct {
    Name string
    Adjective string
    Description string
    Urls []string
}

// Docker, starting from 0.7.x, generates names from notable scientists and hackers.
// Please, for any amazing man that you add to the list, consider adding an equally amazing woman to it, and vice versa.
var persons []Person = []Person{
    Person{
        Name: "agnesi",
        Description: "Maria Gaetana Agnesi. Italian mathematician, philosopher, theologian and humanitarian. She was the first woman to write a mathematics handbook and the first woman appointed as a Mathematics Professor at a University.",
        Urls: []string{"https://en.wikipedia.org/wiki/Maria_Gaetana_Agnesi",},
    },
    Person{
        Name: "albattani",
        Description: "Muhammad ibn J??bir al-???arr??n?? al-Batt??n?? was a founding father of astronomy.",
        Urls: []string{"https://en.wikipedia.org/wiki/Mu%E1%B8%A5ammad_ibn_J%C4%81bir_al-%E1%B8%A4arr%C4%81n%C4%AB_al-Batt%C4%81n%C4%AB",},
    },
    Person{
        Name: "allen",
        Description: "Frances E. Allen, became the first female IBM Fellow in 1989. In 2006, she became the first female recipient of the ACM's Turing Award.",
        Urls: []string{"https://en.wikipedia.org/wiki/Frances_E._Allen",},
    },
    Person{
        Name: "almeida",
        Description: "June Almeida. Scottish virologist who took the first pictures of the rubella virus.",
        Urls: []string{"https://en.wikipedia.org/wiki/June_Almeida",},
    },
    Person{
        Name: "antonelli",
        Description: "Kathleen Antonelli, American computer programmer and one of the six original programmers of the ENIAC.",
        Urls: []string{"https://en.wikipedia.org/wiki/Kathleen_Antonelli",},
    },
    Person{
        Name: "archimedes",
        Description: "Archimedes was a physicist, engineer and mathematician who invented too many things to list them here.",
        Urls: []string{"https://en.wikipedia.org/wiki/Archimedes",},
    },
    Person{
        Name: "ardinghelli",
        Description: "Maria Ardinghelli. Italian translator, mathematician and physicist.",
        Urls: []string{"https://en.wikipedia.org/wiki/Maria_Ardinghelli",},
    },
    Person{
        Name: "aryabhata",
        Description: "Aryabhata. Ancient Indian mathematician-astronomer during 476-550 CE.",
        Urls: []string{"https://en.wikipedia.org/wiki/Aryabhata",},
    },
    Person{
        Name: "austin",
        Description: "Wanda Austin. Wanda Austin is the President and CEO of The Aerospace Corporation, a leading architect for the US security space programs.",
        Urls: []string{"https://en.wikipedia.org/wiki/Wanda_Austin",},
    },
    Person{
        Name: "babbage",
        Description: "Charles Babbage invented the concept of a programmable computer.",
        Urls: []string{"https://en.wikipedia.org/wiki/Charles_Babbage.",},
    },
    Person{
        Name: "banach",
        Description: "Stefan Banach. Polish mathematician, was one of the founders of modern functional analysis.",
        Urls: []string{"https://en.wikipedia.org/wiki/Stefan_Banach",},
    },
    Person{
        Name: "banzai",
        Description: "Buckaroo Banzai and his mentor Dr. Hikita perfected the \"oscillation overthruster\", a device that allows one to pass through solid matter.",
        Urls: []string{"https://en.wikipedia.org/wiki/The_Adventures_of_Buckaroo_Banzai_Across_the_8th_Dimension",},
    },
    Person{
        Name: "bardeen",
        Description: "John Bardeen co-invented the transistor.",
        Urls: []string{"https://en.wikipedia.org/wiki/John_Bardeen",},
    },
    Person{
        Name: "bartik",
        Description: "Jean Bartik, born Betty Jean Jennings, was one of the original programmers for the ENIAC computer.",
        Urls: []string{"https://en.wikipedia.org/wiki/Jean_Bartik",},
    },
    Person{
        Name: "bassi",
        Description: "Laura Bassi, the world's first female professor.",
        Urls: []string{"https://en.wikipedia.org/wiki/Laura_Bassi",},
    },
    Person{
        Name: "beaver",
        Description: "Hugh Beaver, British engineer, founder of the Guinness Book of World Records.",
        Urls: []string{"https://en.wikipedia.org/wiki/Hugh_Beaver",},
    },
    Person{
        Name: "bell",
        Description: "Alexander Graham Bell. An eminent Scottish-born scientist, inventor, engineer and innovator who is credited with inventing the first practical telephone.",
        Urls: []string{"https://en.wikipedia.org/wiki/Alexander_Graham_Bell",},
    },
    Person{
        Name: "benz",
        Description: "Karl Friedrich Benz. A German automobile engineer. Inventor of the first practical motorcar.",
        Urls: []string{"https://en.wikipedia.org/wiki/Karl_Benz",},
    },
    Person{
        Name: "bhabha",
        Description: "Homi J Bhabha. Was an Indian nuclear physicist, founding director, and professor of physics at the Tata Institute of Fundamental Research. Colloquially known as \"father of Indian nuclear programme\".",
        Urls: []string{"https://en.wikipedia.org/wiki/Homi_J._Bhabha",},
    },
    Person{
        Name: "bhaskara",
        Description: "Bhaskara II. Ancient Indian mathematician-astronomer whose work on calculus predates Newton and Leibniz by over half a millennium.",
        Urls: []string{"https://en.wikipedia.org/wiki/Bh%C4%81skara_II#Calculus",},
    },
    Person{
        Name: "black",
        Description: "Sue Black. British computer scientist and campaigner. She has been instrumental in saving Bletchley Park, the site of World War II codebreaking.",
        Urls: []string{"https://en.wikipedia.org/wiki/Sue_Black_(computer_scientist)",},
    },
    Person{
        Name: "blackburn",
        Description: "Elizabeth Helen Blackburn. Australian-American Nobel laureate; best known for co-discovering telomerase.",
        Urls: []string{"https://en.wikipedia.org/wiki/Elizabeth_Blackburn",},
    },
    Person{
        Name: "blackwell",
        Description: "Elizabeth Blackwell. American doctor and first American woman to receive a medical degree.",
        Urls: []string{"https://en.wikipedia.org/wiki/Elizabeth_Blackwell",},
    },
    Person{
        Name: "bohr",
        Description: "Niels Bohr is the father of quantum theory.",
        Urls: []string{"https://en.wikipedia.org/wiki/Niels_Bohr.",},
    },
    Person{
        Name: "booth",
        Description: "Kathleen Booth, she's credited with writing the first assembly language.",
        Urls: []string{"https://en.wikipedia.org/wiki/Kathleen_Booth",},
    },
    Person{
        Name: "borg",
        Description: "Anita Borg. Anita Borg was the founding director of the Institute for Women and Technology (IWT).",
        Urls: []string{"https://en.wikipedia.org/wiki/Anita_Borg",},
    },
    Person{
        Name: "bose",
        Description: "Satyendra Nath Bose. He provided the foundation for Bose???Einstein statistics and the theory of the Bose???Einstein condensate.",
        Urls: []string{"https://en.wikipedia.org/wiki/Satyendra_Nath_Bose",},
    },
    Person{
        Name: "bouman",
        Description: "Katherine Louise Bouman is an imaging scientist and Assistant Professor of Computer Science at the California Institute of Technology. She researches computational methods for imaging, and developed an algorithm that made possible the picture first visualization of a black hole using the Event Horizon Telescope.",
        Urls: []string{"https://en.wikipedia.org/wiki/Katie_Bouman",},
    },
    Person{
        Name: "boyd",
        Description: "Evelyn Boyd Granville. She was one of the first African-American woman to receive a Ph.D. in mathematics; she earned it in 1949 from Yale University.",
        Urls: []string{"https://en.wikipedia.org/wiki/Evelyn_Boyd_Granville",},
    },
    Person{
        Name: "brahmagupta",
        Description: "Brahmagupta. Ancient Indian mathematician during 598-670 CE who gave rules to compute with zero.",
        Urls: []string{"https://en.wikipedia.org/wiki/Brahmagupta#Zero",},
    },
    Person{
        Name: "brattain",
        Description: "Walter Houser Brattain co-invented the transistor.",
        Urls: []string{"https://en.wikipedia.org/wiki/Walter_Houser_Brattain",},
    },
    Person{
        Name: "brown",
        Description: "Emmett Brown invented time travel.",
        Urls: []string{"https://en.wikipedia.org/wiki/Emmett_Brown (thanks Brian Goff)",},
    },
    Person{
        Name: "buck",
        Description: "Linda Brown Buck. American biologist and Nobel laureate best known for her genetic and molecular analyses of the mechanisms of smell.",
        Urls: []string{"https://en.wikipedia.org/wiki/Linda_B._Buck",},
    },
    Person{
        Name: "burnell",
        Description: "Dame Susan Jocelyn Bell Burnell. Northern Irish astrophysicist who discovered radio pulsars and was the first to analyze them.",
        Urls: []string{"https://en.wikipedia.org/wiki/Jocelyn_Bell_Burnell",},
    },
    Person{
        Name: "cannon",
        Description: "Annie Jump Cannon. Pioneering female astronomer who classified hundreds of thousands of stars and created the system we use to understand stars today.",
        Urls: []string{"https://en.wikipedia.org/wiki/Annie_Jump_Cannon",},
    },
    Person{
        Name: "carson",
        Description: "Rachel Carson. American marine biologist and conservationist, her book Silent Spring and other writings are credited with advancing the global environmental movement.",
        Urls: []string{"https://en.wikipedia.org/wiki/Rachel_Carson",},
    },
    Person{
        Name: "cartwright",
        Description: "Dame Mary Lucy Cartwright. British mathematician who was one of the first to study what is now known as chaos theory. Also known for Cartwright's theorem which finds applications in signal processing.",
        Urls: []string{"https://en.wikipedia.org/wiki/Mary_Cartwright",},
    },
    Person{
        Name: "carver",
        Description: "George Washington Carver. American agricultural scientist and inventor. He was the most prominent black scientist of the early 20th century.",
        Urls: []string{"https://en.wikipedia.org/wiki/George_Washington_Carver",},
    },
    Person{
        Name: "cerf",
        Description: "Vinton Gray Cerf. American Internet pioneer, recognized as one of \"the fathers of the Internet\". With Robert Elliot Kahn, he designed TCP and IP, the primary data communication protocols of the Internet and other computer networks.",
        Urls: []string{"https://en.wikipedia.org/wiki/Vint_Cerf",},
    },
    Person{
        Name: "chandrasekhar",
        Description: "Subrahmanyan Chandrasekhar. Astrophysicist known for his mathematical theory on different stages and evolution in structures of the stars. He has won nobel prize for physics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Subrahmanyan_Chandrasekhar",},
    },
    Person{
        Name: "chaplygin",
        Description: "Sergey Alexeyevich Chaplygin (Russian: ?????????????? ?????????????????????? ??????????????????; April 5, 1869 ??? October 8, 1942) was a Russian and Soviet physicist, mathematician, and mechanical engineer. He is known for mathematical formulas such as Chaplygin's equation and for a hypothetical substance in cosmology called Chaplygin gas, named after him.",
        Urls: []string{"https://en.wikipedia.org/wiki/Sergey_Chaplygin",},
    },
    Person{
        Name: "chatelet",
        Description: "??milie du Ch??telet. French natural philosopher, mathematician, physicist, and author during the early 1730s, known for her translation of and commentary on Isaac Newton's book Principia containing basic laws of physics.",
        Urls: []string{"https://en.wikipedia.org/wiki/%C3%89milie_du_Ch%C3%A2telet",},
    },
    Person{
        Name: "chatterjee",
        Description: "Asima Chatterjee was an Indian organic chemist noted for her research on vinca alkaloids, development of drugs for treatment of epilepsy and malaria.",
        Urls: []string{"https://en.wikipedia.org/wiki/Asima_Chatterjee",},
    },
    Person{
        Name: "chaum",
        Description: "David Lee Chaum. American computer scientist and cryptographer. Known for his seminal contributions in the field of anonymous communication.",
        Urls: []string{"https://en.wikipedia.org/wiki/David_Chaum",},
    },
    Person{
        Name: "chebyshev",
        Description: "Pafnuty Chebyshev. Russian mathematician. He is known fo his works on probability, statistics, mechanics, analytical geometry and number theory.",
        Urls: []string{"https://en.wikipedia.org/wiki/Pafnuty_Chebyshev",},
    },
    Person{
        Name: "clarke",
        Description: "Joan Clarke. Bletchley Park code breaker during the Second World War who pioneered techniques that remained top secret for decades. Also an accomplished numismatist.",
        Urls: []string{"https://en.wikipedia.org/wiki/Joan_Clarke",},
    },
    Person{
        Name: "cohen",
        Description: "Bram Cohen. American computer programmer and author of the BitTorrent peer-to-peer protocol.",
        Urls: []string{"https://en.wikipedia.org/wiki/Bram_Cohen",},
    },
    Person{
        Name: "colden",
        Description: "Jane Colden. American botanist widely considered the first female American botanist.",
        Urls: []string{"https://en.wikipedia.org/wiki/Jane_Colden",},
    },
    Person{
        Name: "cori",
        Description: "Gerty Theresa Cori. American biochemist who became the third woman???and first American woman???to win a Nobel Prize in science, and the first woman to be awarded the Nobel Prize in Physiology or Medicine. Cori was born in Prague.",
        Urls: []string{"https://en.wikipedia.org/wiki/Gerty_Cori",},
    },
    Person{
        Name: "cray",
        Description: "Seymour Roger Cray was an American electrical engineer and supercomputer architect who designed a series of computers that were the fastest in the world for decades.",
        Urls: []string{"https://en.wikipedia.org/wiki/Seymour_Cray",},
    },
    Person{
        Name: "curran",
        Description: "This entry reflects a husband and wife team who worked together: Joan Curran was a Welsh scientist who developed radar and invented chaff, a radar countermeasure. Samuel Curran was an Irish physicist who worked alongside his wife during WWII and invented the proximity fuse.",
        Urls: []string{"https://en.wikipedia.org/wiki/Joan_Curran", "https://en.wikipedia.org/wiki/Samuel_Curran",},
    },
    Person{
        Name: "curie",
        Description: "Marie Curie discovered radioactivity.",
        Urls: []string{"https://en.wikipedia.org/wiki/Marie_Curie.",},
    },
    Person{
        Name: "darwin",
        Description: "Charles Darwin established the principles of natural evolution.",
        Urls: []string{"https://en.wikipedia.org/wiki/Charles_Darwin.",},
    },
    Person{
        Name: "davinci",
        Description: "Leonardo Da Vinci invented too many things to list here.",
        Urls: []string{"https://en.wikipedia.org/wiki/Leonardo_da_Vinci.",},
    },
    Person{
        Name: "dewdney",
        Description: "A. K. (Alexander Keewatin) Dewdney, Canadian mathematician, computer scientist, author and filmmaker. Contributor to Scientific American's \"Computer Recreations\" from 1984 to 1991. Author of Core War (program), The Planiverse, The Armchair Universe, The Magic Machine, The New Turing Omnibus, and more.",
        Urls: []string{"https://en.wikipedia.org/wiki/Alexander_Dewdney",},
    },
    Person{
        Name: "dhawan",
        Description: "Satish Dhawan. Indian mathematician and aerospace engineer, known for leading the successful and indigenous development of the Indian space programme.",
        Urls: []string{"https://en.wikipedia.org/wiki/Satish_Dhawan",},
    },
    Person{
        Name: "diffie",
        Description: "Bailey Whitfield Diffie. American cryptographer and one of the pioneers of public-key cryptography.",
        Urls: []string{"https://en.wikipedia.org/wiki/Whitfield_Diffie",},
    },
    Person{
        Name: "dijkstra",
        Description: "Edsger Wybe Dijkstra was a Dutch computer scientist and mathematical scientist.",
        Urls: []string{"https://en.wikipedia.org/wiki/Edsger_W._Dijkstra.",},
    },
    Person{
        Name: "dirac",
        Description: "Paul Adrien Maurice Dirac. English theoretical physicist who made fundamental contributions to the early development of both quantum mechanics and quantum electrodynamics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Paul_Dirac",},
    },
    Person{
        Name: "driscoll",
        Description: "Agnes Meyer Driscoll. American cryptanalyst during World Wars I and II who successfully cryptanalysed a number of Japanese ciphers. She was also the co-developer of one of the cipher machines of the US Navy, the CM.",
        Urls: []string{"https://en.wikipedia.org/wiki/Agnes_Meyer_Driscoll",},
    },
    Person{
        Name: "dubinsky",
        Description: "Donna Dubinsky. Played an integral role in the development of personal digital assistants (PDAs) serving as CEO of Palm, Inc. and co-founding Handspring.",
        Urls: []string{"https://en.wikipedia.org/wiki/Donna_Dubinsky",},
    },
    Person{
        Name: "easley",
        Description: "Annie Easley. She was a leading member of the team which developed software for the Centaur rocket stage and one of the first African-Americans in her field.",
        Urls: []string{"https://en.wikipedia.org/wiki/Annie_Easley",},
    },
    Person{
        Name: "edison",
        Description: "Thomas Alva Edison, prolific inventor.",
        Urls: []string{"https://en.wikipedia.org/wiki/Thomas_Edison",},
    },
    Person{
        Name: "einstein",
        Description: "Albert Einstein invented the general theory of relativity.",
        Urls: []string{"https://en.wikipedia.org/wiki/Albert_Einstein",},
    },
    Person{
        Name: "elbakyan",
        Description: "Alexandra Asanovna Elbakyan (Russian: ?????????????????????? ?????????????????? ????????????????) is a Kazakhstani graduate student, computer programmer, internet pirate in hiding, and the creator of the site Sci-Hub. Nature has listed her in 2016 in the top ten people that mattered in science, and Ars Technica has compared her to Aaron Swartz.",
        Urls: []string{"https://en.wikipedia.org/wiki/Alexandra_Elbakyan",},
    },
    Person{
        Name: "elgamal",
        Description: "Taher A. ElGamal. Egyptian cryptographer best known for the ElGamal discrete log cryptosystem and the ElGamal digital signature scheme.",
        Urls: []string{"https://en.wikipedia.org/wiki/Taher_Elgamal",},
    },
    Person{
        Name: "elion",
        Description: "Gertrude Elion. American biochemist, pharmacologist and the 1988 recipient of the Nobel Prize in Medicine.",
        Urls: []string{"https://en.wikipedia.org/wiki/Gertrude_Elion",},
    },
    Person{
        Name: "ellis",
        Description: "James Henry Ellis. British engineer and cryptographer employed by the GCHQ. Best known for conceiving for the first time, the idea of public-key cryptography.",
        Urls: []string{"https://en.wikipedia.org/wiki/James_H._Ellis",},
    },
    Person{
        Name: "engelbart",
        Description: "Douglas Engelbart gave the mother of all demos.",
        Urls: []string{"https://en.wikipedia.org/wiki/Douglas_Engelbart",},
    },
    Person{
        Name: "euclid",
        Description: "Euclid invented geometry.",
        Urls: []string{"https://en.wikipedia.org/wiki/Euclid",},
    },
    Person{
        Name: "euler",
        Description: "Leonhard Euler invented large parts of modern mathematics.",
        Urls: []string{"https://de.wikipedia.org/wiki/Leonhard_Euler",},
    },
    Person{
        Name: "faraday",
        Description: "Michael Faraday. British scientist who contributed to the study of electromagnetism and electrochemistry.",
        Urls: []string{"https://en.wikipedia.org/wiki/Michael_Faraday",},
    },
    Person{
        Name: "feistel",
        Description: "Horst Feistel. German-born American cryptographer who was one of the earliest non-government researchers to study the design and theory of block ciphers. Co-developer of DES and Lucifer. Feistel networks, a symmetric structure used in the construction of block ciphers are named after him.",
        Urls: []string{"https://en.wikipedia.org/wiki/Horst_Feistel",},
    },
    Person{
        Name: "fermat",
        Description: "Pierre de Fermat pioneered several aspects of modern mathematics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Pierre_de_Fermat",},
    },
    Person{
        Name: "fermi",
        Description: "Enrico Fermi invented the first nuclear reactor.",
        Urls: []string{"https://en.wikipedia.org/wiki/Enrico_Fermi.",},
    },
    Person{
        Name: "feynman",
        Description: "Richard Feynman was a key contributor to quantum mechanics and particle physics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Richard_Feynman",},
    },
    Person{
        Name: "franklin",
        Description: "Benjamin Franklin is famous for his experiments in electricity and the invention of the lightning rod.",
        Urls: []string{},
    },
    Person{
        Name: "gagarin",
        Description: "Yuri Alekseyevich Gagarin. Soviet pilot and cosmonaut, best known as the first human to journey into outer space.",
        Urls: []string{"https://en.wikipedia.org/wiki/Yuri_Gagarin",},
    },
    Person{
        Name: "galileo",
        Description: "Galileo was a founding father of modern astronomy, and faced politics and obscurantism to establish scientific truth.",
        Urls: []string{"https://en.wikipedia.org/wiki/Galileo_Galilei",},
    },
    Person{
        Name: "galois",
        Description: "??variste Galois. French mathematician whose work laid the foundations of Galois theory and group theory, two major branches of abstract algebra, and the subfield of Galois connections, all while still in his late teens.",
        Urls: []string{"https://en.wikipedia.org/wiki/%C3%89variste_Galois",},
    },
    Person{
        Name: "ganguly",
        Description: "Kadambini Ganguly. Indian physician, known for being the first South Asian female physician, trained in western medicine, to graduate in South Asia.",
        Urls: []string{"https://en.wikipedia.org/wiki/Kadambini_Ganguly",},
    },
    Person{
        Name: "gates",
        Description: "William Henry \"Bill\" Gates III is an American business magnate, philanthropist, investor, computer programmer, and inventor.",
        Urls: []string{"https://en.wikipedia.org/wiki/Bill_Gates",},
    },
    Person{
        Name: "gauss",
        Description: "Johann Carl Friedrich Gauss. German mathematician who made significant contributions to many fields, including number theory, algebra, statistics, analysis, differential geometry, geodesy, geophysics, mechanics, electrostatics, magnetic fields, astronomy, matrix theory, and optics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Carl_Friedrich_Gauss",},
    },
    Person{
        Name: "germain",
        Description: "Marie-Sophie Germain. French mathematician, physicist and philosopher. Known for her work on elasticity theory, number theory and philosophy.",
        Urls: []string{"https://en.wikipedia.org/wiki/Sophie_Germain",},
    },
    Person{
        Name: "goldberg",
        Description: "Adele Goldberg, was one of the designers and developers of the Smalltalk language.",
        Urls: []string{"https://en.wikipedia.org/wiki/Adele_Goldberg_(computer_scientist)",},
    },
    Person{
        Name: "goldstine",
        Description: "Adele Goldstine, born Adele Katz, wrote the complete technical description for the first electronic digital computer, ENIAC.",
        Urls: []string{"https://en.wikipedia.org/wiki/Adele_Goldstine",},
    },
    Person{
        Name: "goldwasser",
        Description: "Shafi Goldwasser is a computer scientist known for creating theoretical foundations of modern cryptography. Winner of 2012 ACM Turing Award.",
        Urls: []string{"https://en.wikipedia.org/wiki/Shafi_Goldwasser",},
    },
    Person{
        Name: "golick",
        Description: "James Golick, all around gangster.",
        Urls: []string{},
    },
    Person{
        Name: "goodall",
        Description: "Jane Goodall. British primatologist, ethologist, and anthropologist who is considered to be the world's foremost expert on chimpanzees.",
        Urls: []string{"https://en.wikipedia.org/wiki/Jane_Goodall",},
    },
    Person{
        Name: "gould",
        Description: "Stephen Jay Gould was was an American paleontologist, evolutionary biologist, and historian of science. He is most famous for the theory of punctuated equilibrium.",
        Urls: []string{"https://en.wikipedia.org/wiki/Stephen_Jay_Gould",},
    },
    Person{
        Name: "greider",
        Description: "Carolyn Widney Greider. American molecular biologist and joint winner of the 2009 Nobel Prize for Physiology or Medicine for the discovery of telomerase.",
        Urls: []string{"https://en.wikipedia.org/wiki/Carol_W._Greider",},
    },
    Person{
        Name: "grothendieck",
        Description: "Alexander Grothendieck. German-born French mathematician who became a leading figure in the creation of modern algebraic geometry.",
        Urls: []string{"https://en.wikipedia.org/wiki/Alexander_Grothendieck",},
    },
    Person{
        Name: "haibt",
        Description: "Lois Haibt. American computer scientist, part of the team at IBM that developed FORTRAN.",
        Urls: []string{"https://en.wikipedia.org/wiki/Lois_Haibt",},
    },
    Person{
        Name: "hamilton",
        Description: "Margaret Hamilton. Director of the Software Engineering Division of the MIT Instrumentation Laboratory, which developed on-board flight software for the Apollo space program.",
        Urls: []string{"https://en.wikipedia.org/wiki/Margaret_Hamilton_(scientist)",},
    },
    Person{
        Name: "haslett",
        Description: "Caroline Harriet Haslett. English electrical engineer, electricity industry administrator and champion of women's rights. Co-author of British Standard 1363 that specifies AC power plugs and sockets used across the United Kingdom (which is widely considered as one of the safest designs).",
        Urls: []string{"https://en.wikipedia.org/wiki/Caroline_Haslett",},
    },
    Person{
        Name: "hawking",
        Description: "Stephen Hawking pioneered the field of cosmology by combining general relativity and quantum mechanics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Stephen_Hawking",},
    },
    Person{
        Name: "hellman",
        Description: "Martin Edward Hellman. American cryptologist, best known for his invention of public-key cryptography in co-operation with Whitfield Diffie and Ralph Merkle.",
        Urls: []string{"https://en.wikipedia.org/wiki/Martin_Hellman",},
    },
    Person{
        Name: "heisenberg",
        Description: "Werner Heisenberg was a founding father of quantum mechanics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Werner_Heisenberg",},
    },
    Person{
        Name: "hermann",
        Description: "Grete Hermann was a German philosopher noted for her philosophical work on the foundations of quantum mechanics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Grete_Hermann",},
    },
    Person{
        Name: "herschel",
        Description: "Caroline Lucretia Herschel. German astronomer and discoverer of several comets.",
        Urls: []string{"https://en.wikipedia.org/wiki/Caroline_Herschel",},
    },
    Person{
        Name: "hertz",
        Description: "Heinrich Rudolf Hertz. German physicist who first conclusively proved the existence of the electromagnetic waves.",
        Urls: []string{"https://en.wikipedia.org/wiki/Heinrich_Hertz",},
    },
    Person{
        Name: "heyrovsky",
        Description: "Jaroslav Heyrovsk?? was the inventor of the polarographic method, father of the electroanalytical method, and recipient of the Nobel Prize in 1959. His main field of work was polarography.",
        Urls: []string{"https://en.wikipedia.org/wiki/Jaroslav_Heyrovsk%C3%BD",},
    },
    Person{
        Name: "hodgkin",
        Description: "Dorothy Hodgkin was a British biochemist, credited with the development of protein crystallography. She was awarded the Nobel Prize in Chemistry in 1964.",
        Urls: []string{"https://en.wikipedia.org/wiki/Dorothy_Hodgkin",},
    },
    Person{
        Name: "hofstadter",
        Description: "Douglas R. Hofstadter is an American professor of cognitive science and author of the Pulitzer Prize and American Book Award-winning work Goedel, Escher, Bach: An Eternal Golden Braid in 1979. A mind-bending work which coined Hofstadter's Law: \"It always takes longer than you expect, even when you take into account Hofstadter's Law.\".",
        Urls: []string{"https://en.wikipedia.org/wiki/Douglas_Hofstadter",},
    },
    Person{
        Name: "hoover",
        Description: "Erna Schneider Hoover revolutionized modern communication by inventing a computerized telephone switching method.",
        Urls: []string{"https://en.wikipedia.org/wiki/Erna_Schneider_Hoover",},
    },
    Person{
        Name: "hopper",
        Description: "Grace Hopper developed the first compiler for a computer programming language and is credited with popularizing the term \"debugging\" for fixing computer glitches.",
        Urls: []string{"https://en.wikipedia.org/wiki/Grace_Hopper",},
    },
    Person{
        Name: "hugle",
        Description: "Frances Hugle, she was an American scientist, engineer, and inventor who contributed to the understanding of semiconductors, integrated circuitry, and the unique electrical principles of microscopic materials.",
        Urls: []string{"https://en.wikipedia.org/wiki/Frances_Hugle",},
    },
    Person{
        Name: "hypatia",
        Description: "Hypatia. Greek Alexandrine Neoplatonist philosopher in Egypt who was one of the earliest mothers of mathematics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Hypatia",},
    },
    Person{
        Name: "ishizaka",
        Description: "Teruko Ishizaka. Japanese scientist and immunologist who co-discovered the antibody class Immunoglobulin E.",
        Urls: []string{"https://en.wikipedia.org/wiki/Teruko_Ishizaka",},
    },
    Person{
        Name: "jackson",
        Description: "Mary Jackson, American mathematician and aerospace engineer who earned the highest title within NASA's engineering department.",
        Urls: []string{"https://en.wikipedia.org/wiki/Mary_Jackson_(engineer)",},
    },
    Person{
        Name: "jang",
        Description: "Yeong-Sil Jang was a Korean scientist and astronomer during the Joseon Dynasty; he invented the first metal printing press and water gauge.",
        Urls: []string{"https://en.wikipedia.org/wiki/Jang_Yeong-sil",},
    },
    Person{
        Name: "jemison",
        Description: "Mae Carol Jemison, is an American engineer, physician, and former NASA astronaut. She became the first black woman to travel in space when she served as a mission specialist aboard the Space Shuttle Endeavour.",
        Urls: []string{"https://en.wikipedia.org/wiki/Mae_Jemison",},
    },
    Person{
        Name: "jennings",
        Description: "Betty Jennings. One of the original programmers of the ENIAC.",
        Urls: []string{"https://en.wikipedia.org/wiki/ENIAC", "https://en.wikipedia.org/wiki/Jean_Bartik",},
    },
    Person{
        Name: "jepsen",
        Description: "Mary Lou Jepsen, was the founder and chief technology officer of One Laptop Per Child (OLPC), and the founder of Pixel Qi.",
        Urls: []string{"https://en.wikipedia.org/wiki/Mary_Lou_Jepsen",},
    },
    Person{
        Name: "johnson",
        Description: "Katherine Coleman Goble Johnson. American physicist and mathematician contributed to the NASA.",
        Urls: []string{"https://en.wikipedia.org/wiki/Katherine_Johnson",},
    },
    Person{
        Name: "joliot",
        Description: "Ir??ne Joliot-Curie. French scientist who was awarded the Nobel Prize for Chemistry in 1935. Daughter of Marie and Pierre Curie.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ir%C3%A8ne_Joliot-Curie",},
    },
    Person{
        Name: "jones",
        Description: "Karen Sp??rck Jones came up with the concept of inverse document frequency, which is used in most search engines today.",
        Urls: []string{"https://en.wikipedia.org/wiki/Karen_Sp%C3%A4rck_Jones",},
    },
    Person{
        Name: "kalam",
        Description: "A. P. J. Abdul Kalam, is an Indian scientist aka Missile Man of India for his work on the development of ballistic missile and launch vehicle technology.",
        Urls: []string{"https://en.wikipedia.org/wiki/A._P._J._Abdul_Kalam",},
    },
    Person{
        Name: "kapitsa",
        Description: "Sergey Petrovich Kapitsa (Russian: ?????????????? ?????????????????? ??????????????; 14 February 1928 ??? 14 August 2012) was a Russian physicist and demographer. He was best known as host of the popular and long-running Russian scientific TV show, Evident, but Incredible. His father was the Nobel laureate Soviet-era physicist Pyotr Kapitsa, and his brother was the geographer and Antarctic explorer Andrey Kapitsa.",
        Urls: []string{"https://en.wikipedia.org/wiki/Sergey_Kapitsa",},
    },
    Person{
        Name: "kare",
        Description: "Susan Kare, created the icons and many of the interface elements for the original Apple Macintosh in the 1980s, and was an original employee of NeXT, working as the Creative Director.",
        Urls: []string{"https://en.wikipedia.org/wiki/Susan_Kare",},
    },
    Person{
        Name: "keldysh",
        Description: "Mstislav Keldysh. A Soviet scientist in the field of mathematics and mechanics, academician of the USSR Academy of Sciences (1946), President of the USSR Academy of Sciences (1961???1975), three times Hero of Socialist Labor (1956, 1961, 1971), fellow of the Royal Society of Edinburgh (1968).",
        Urls: []string{"https://en.wikipedia.org/wiki/Mstislav_Keldysh",},
    },
    Person{
        Name: "keller",
        Description: "Mary Kenneth Keller, Sister Mary Kenneth Keller became the first American woman to earn a PhD in Computer Science in 1965.",
        Urls: []string{"https://en.wikipedia.org/wiki/Mary_Kenneth_Keller",},
    },
    Person{
        Name: "kepler",
        Description: "Johannes Kepler, German astronomer known for his three laws of planetary motion.",
        Urls: []string{"https://en.wikipedia.org/wiki/Johannes_Kepler",},
    },
    Person{
        Name: "khayyam",
        Description: "Omar Khayyam. Persian mathematician, astronomer and poet. Known for his work on the classification and solution of cubic equations, for his contribution to the understanding of Euclid's fifth postulate and for computing the length of a year very accurately.",
        Urls: []string{"https://en.wikipedia.org/wiki/Omar_Khayyam",},
    },
    Person{
        Name: "khorana",
        Description: "Har Gobind Khorana. Indian-American biochemist who shared the 1968 Nobel Prize for Physiology.",
        Urls: []string{"https://en.wikipedia.org/wiki/Har_Gobind_Khorana",},
    },
    Person{
        Name: "kilby",
        Description: "Jack Kilby invented silicon integrated circuits and gave Silicon Valley its name.",
        Urls: []string{"https://en.wikipedia.org/wiki/Jack_Kilby",},
    },
    Person{
        Name: "kirch",
        Description: "Maria Kirch. German astronomer and first woman to discover a comet.",
        Urls: []string{"https://en.wikipedia.org/wiki/Maria_Margarethe_Kirch",},
    },
    Person{
        Name: "knuth",
        Description: "Donald Knuth. American computer scientist, author of \"The Art of Computer Programming\" and creator of the TeX typesetting system.",
        Urls: []string{"https://en.wikipedia.org/wiki/Donald_Knuth",},
    },
    Person{
        Name: "kowalevski",
        Description: "Sophie Kowalevski. Russian mathematician responsible for important original contributions to analysis, differential equations and mechanics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Sofia_Kovalevskaya",},
    },
    Person{
        Name: "lalande",
        Description: "Marie-Jeanne de Lalande. French astronomer, mathematician and cataloguer of stars.",
        Urls: []string{"https://en.wikipedia.org/wiki/Marie-Jeanne_de_Lalande",},
    },
    Person{
        Name: "lamarr",
        Description: "Hedy Lamarr. Actress and inventor. The principles of her work are now incorporated into modern Wi-Fi, CDMA and Bluetooth technology.",
        Urls: []string{"https://en.wikipedia.org/wiki/Hedy_Lamarr",},
    },
    Person{
        Name: "lamport",
        Description: "Leslie B. Lamport. American computer scientist. Lamport is best known for his seminal work in distributed systems and was the winner of the 2013 Turing Award.",
        Urls: []string{"https://en.wikipedia.org/wiki/Leslie_Lamport",},
    },
    Person{
        Name: "leakey",
        Description: "Mary Leakey. British paleoanthropologist who discovered the first fossilized Proconsul skull.",
        Urls: []string{"https://en.wikipedia.org/wiki/Mary_Leakey",},
    },
    Person{
        Name: "leavitt",
        Description: "Henrietta Swan Leavitt. She was an American astronomer who discovered the relation between the luminosity and the period of Cepheid variable stars.",
        Urls: []string{"https://en.wikipedia.org/wiki/Henrietta_Swan_Leavitt",},
    },
    Person{
        Name: "lederberg",
        Description: "Esther Miriam Zimmer Lederberg. American microbiologist and a pioneer of bacterial genetics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Esther_Lederberg",},
    },
    Person{
        Name: "lehmann",
        Description: "Inge Lehmann. Danish seismologist and geophysicist. Known for discovering in 1936 that the Earth has a solid inner core inside a molten outer core.",
        Urls: []string{"https://en.wikipedia.org/wiki/Inge_Lehmann",},
    },
    Person{
        Name: "lewin",
        Description: "Daniel Lewin. Mathematician, Akamai co-founder, soldier, 9/11 victim-- Developed optimization techniques for routing traffic on the internet. Died attempting to stop the 9-11 hijackers.",
        Urls: []string{"https://en.wikipedia.org/wiki/Daniel_Lewin",},
    },
    Person{
        Name: "lichterman",
        Description: "Ruth Lichterman. One of the original programmers of the ENIAC.",
        Urls: []string{"https://en.wikipedia.org/wiki/ENIAC", "https://en.wikipedia.org/wiki/Ruth_Teitelbaum",},
    },
    Person{
        Name: "liskov",
        Description: "Barbara Liskov, co-developed the Liskov substitution principle. Liskov was also the winner of the Turing Prize in 2008.",
        Urls: []string{"https://en.wikipedia.org/wiki/Barbara_Liskov",},
    },
    Person{
        Name: "lovelace",
        Description: "Ada Lovelace invented the first algorithm.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ada_Lovelace (thanks James Turnbull)",},
    },
    Person{
        Name: "lumiere",
        Description: "Auguste and Louis Lumi??re, the first filmmakers in history.",
        Urls: []string{"https://en.wikipedia.org/wiki/Auguste_and_Louis_Lumi%C3%A8re",},
    },
    Person{
        Name: "mahavira",
        Description: "Mahavira. Ancient Indian mathematician during 9th century AD who discovered basic algebraic identities.",
        Urls: []string{"https://en.wikipedia.org/wiki/Mah%C4%81v%C4%ABra_(mathematician)",},
    },
    Person{
        Name: "margulis",
        Description: "Lynn Margulis (b. Lynn Petra Alexander). An American evolutionary theorist and biologist, science author, educator, and popularizer, and was the primary modern proponent for the significance of symbiosis in evolution.",
        Urls: []string{"https://en.wikipedia.org/wiki/Lynn_Margulis",},
    },
    Person{
        Name: "matsumoto",
        Description: "Yukihiro Matsumoto. Japanese computer scientist and software programmer best known as the chief designer of the Ruby programming language.",
        Urls: []string{"https://en.wikipedia.org/wiki/Yukihiro_Matsumoto",},
    },
    Person{
        Name: "maxwell",
        Description: "James Clerk Maxwell. Scottish physicist, best known for his formulation of electromagnetic theory.",
        Urls: []string{"https://en.wikipedia.org/wiki/James_Clerk_Maxwell",},
    },
    Person{
        Name: "mayer",
        Description: "Maria Mayer. American theoretical physicist and Nobel laureate in Physics for proposing the nuclear shell model of the atomic nucleus.",
        Urls: []string{"https://en.wikipedia.org/wiki/Maria_Mayer",},
    },
    Person{
        Name: "mccarthy",
        Description: "John McCarthy invented LISP.",
        Urls: []string{"https://en.wikipedia.org/wiki/John_McCarthy_(computer_scientist)",},
    },
    Person{
        Name: "mcclintock",
        Description: "Barbara McClintock. A distinguished American cytogeneticist, 1983 Nobel Laureate in Physiology or Medicine for discovering transposons.",
        Urls: []string{"https://en.wikipedia.org/wiki/Barbara_McClintock",},
    },
    Person{
        Name: "mclaren",
        Description: "Anne Laura Dorinthea McLaren. British developmental biologist whose work helped lead to human in-vitro fertilisation.",
        Urls: []string{"https://en.wikipedia.org/wiki/Anne_McLaren",},
    },
    Person{
        Name: "mclean",
        Description: "Malcolm McLean invented the modern shipping container.",
        Urls: []string{"https://en.wikipedia.org/wiki/Malcom_McLean",},
    },
    Person{
        Name: "mcnulty",
        Description: "Kay McNulty. One of the original programmers of the ENIAC.",
        Urls: []string{"https://en.wikipedia.org/wiki/ENIAC", "https://en.wikipedia.org/wiki/Kathleen_Antonelli",},
    },
    Person{
        Name: "mendel",
        Description: "Gregor Johann Mendel. Czech scientist and founder of genetics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Gregor_Mendel",},
    },
    Person{
        Name: "mendeleev",
        Description: "Dmitri Mendeleev. A chemist and inventor. He formulated the Periodic Law, created a farsighted version of the periodic table of elements, and used it to correct the properties of some already discovered elements and also to predict the properties of eight elements yet to be discovered.",
        Urls: []string{"https://en.wikipedia.org/wiki/Dmitri_Mendeleev",},
    },
    Person{
        Name: "meitner",
        Description: "Lise Meitner. Austrian/Swedish physicist who was involved in the discovery of nuclear fission. The element meitnerium is named after her.",
        Urls: []string{"https://en.wikipedia.org/wiki/Lise_Meitner",},
    },
    Person{
        Name: "meninsky",
        Description: "Carla Meninsky, was the game designer and programmer for Atari 2600 games Dodge 'Em and Warlords.",
        Urls: []string{"https://en.wikipedia.org/wiki/Carla_Meninsky",},
    },
    Person{
        Name: "merkle",
        Description: "Ralph C. Merkle. American computer scientist, known for devising Merkle's puzzles. One of the very first schemes for public-key cryptography. Also, inventor of Merkle trees and co-inventor of the Merkle-Damg??rd construction for building collision-resistant cryptographic hash functions and the Merkle-Hellman knapsack cryptosystem.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ralph_Merkle",},
    },
    Person{
        Name: "mestorf",
        Description: "Johanna Mestorf. German prehistoric archaeologist and first female museum director in Germany.",
        Urls: []string{"https://en.wikipedia.org/wiki/Johanna_Mestorf",},
    },
    Person{
        Name: "mirzakhani",
        Description: "Maryam Mirzakhani. An Iranian mathematician and the first woman to win the Fields Medal.",
        Urls: []string{"https://en.wikipedia.org/wiki/Maryam_Mirzakhani",},
    },
    Person{
        Name: "montalcini",
        Description: "Rita Levi-Montalcini. Won Nobel Prize in Physiology or Medicine jointly with colleague Stanley Cohen for the discovery of nerve growth factor.",
        Urls: []string{"https://en.wikipedia.org/wiki/Rita_Levi-Montalcini",},
    },
    Person{
        Name: "moore",
        Description: "Gordon Earle Moore. American engineer, Silicon Valley founding father, author of Moore's law.",
        Urls: []string{"https://en.wikipedia.org/wiki/Gordon_Moore",},
    },
    Person{
        Name: "morse",
        Description: "Samuel Morse, contributed to the invention of a single-wire telegraph system based on European telegraphs and was a co-developer of the Morse code.",
        Urls: []string{"https://en.wikipedia.org/wiki/Samuel_Morse",},
    },
    Person{
        Name: "murdock",
        Description: "Ian Murdock, founder of the Debian project.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ian_Murdock",},
    },
    Person{
        Name: "moser",
        Description: "May-Britt Moser. Nobel prize winner neuroscientist who contributed to the discovery of grid cells in the brain.",
        Urls: []string{"https://en.wikipedia.org/wiki/May-Britt_Moser",},
    },
    Person{
        Name: "napier",
        Description: "John Napier of Merchiston. Scottish landowner known as an astronomer, mathematician and physicist. Best known for his discovery of logarithms.",
        Urls: []string{"https://en.wikipedia.org/wiki/John_Napier",},
    },
    Person{
        Name: "nash",
        Description: "John Forbes Nash, Jr.. American mathematician who made fundamental contributions to game theory, differential geometry, and the study of partial differential equations.",
        Urls: []string{"https://en.wikipedia.org/wiki/John_Forbes_Nash_Jr.",},
    },
    Person{
        Name: "neumann",
        Description: "John von Neumann, today's computer architectures are based on the von Neumann architecture.",
        Urls: []string{"https://en.wikipedia.org/wiki/Von_Neumann_architecture",},
    },
    Person{
        Name: "newton",
        Description: "Isaac Newton invented classic mechanics and modern optics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Isaac_Newton",},
    },
    Person{
        Name: "nightingale",
        Description: "Florence Nightingale, more prominently known as a nurse, was also the first female member of the Royal Statistical Society and a pioneer in statistical graphics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Florence_Nightingale#Statistics_and_sanitary_reform",},
    },
    Person{
        Name: "nobel",
        Description: "Alfred Nobel. A Swedish chemist, engineer, innovator, and armaments manufacturer (inventor of dynamite).",
        Urls: []string{"https://en.wikipedia.org/wiki/Alfred_Nobel",},
    },
    Person{
        Name: "noether",
        Description: "Emmy Noether, German mathematician. Noether's Theorem is named after her.",
        Urls: []string{"https://en.wikipedia.org/wiki/Emmy_Noether",},
    },
    Person{
        Name: "northcutt",
        Description: "Poppy Northcutt. Poppy Northcutt was the first woman to work as part of NASA???s Mission Control.",
        Urls: []string{"http://www.businessinsider.com/poppy-northcutt-helped-apollo-astronauts-2014-12?op=1",},
    },
    Person{
        Name: "noyce",
        Description: "Robert Noyce invented silicon integrated circuits and gave Silicon Valley its name.",
        Urls: []string{"https://en.wikipedia.org/wiki/Robert_Noyce",},
    },
    Person{
        Name: "panini",
        Description: "Panini. Ancient Indian linguist and grammarian from 4th century CE who worked on the world's first formal system.",
        Urls: []string{"https://en.wikipedia.org/wiki/P%C4%81%E1%B9%87ini#Comparison_with_modern_formal_systems",},
    },
    Person{
        Name: "pare",
        Description: "Ambroise Pare invented modern surgery.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ambroise_Par%C3%A9",},
    },
    Person{
        Name: "pascal",
        Description: "Blaise Pascal, French mathematician, physicist, and inventor.",
        Urls: []string{"https://en.wikipedia.org/wiki/Blaise_Pascal",},
    },
    Person{
        Name: "pasteur",
        Description: "Louis Pasteur discovered vaccination, fermentation and pasteurization.",
        Urls: []string{"https://en.wikipedia.org/wiki/Louis_Pasteur.",},
    },
    Person{
        Name: "payne",
        Description: "Cecilia Payne-Gaposchkin was an astronomer and astrophysicist who, in 1925, proposed in her Ph.D. thesis an explanation for the composition of stars in terms of the relative abundances of hydrogen and helium.",
        Urls: []string{"https://en.wikipedia.org/wiki/Cecilia_Payne-Gaposchkin",},
    },
    Person{
        Name: "perlman",
        Description: "Radia Perlman is a software designer and network engineer and most famous for her invention of the spanning-tree protocol (STP).",
        Urls: []string{"https://en.wikipedia.org/wiki/Radia_Perlman",},
    },
    Person{
        Name: "pike",
        Description: "Rob Pike was a key contributor to Unix, Plan 9, the X graphic system, utf-8, and the Go programming language.",
        Urls: []string{"https://en.wikipedia.org/wiki/Rob_Pike",},
    },
    Person{
        Name: "poincare",
        Description: "Henri Poincar?? made fundamental contributions in several fields of mathematics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Henri_Poincar%C3%A9",},
    },
    Person{
        Name: "poitras",
        Description: "Laura Poitras is a director and producer whose work, made possible by open source crypto tools, advances the causes of truth and freedom of information by reporting disclosures by whistleblowers such as Edward Snowden.",
        Urls: []string{"https://en.wikipedia.org/wiki/Laura_Poitras",},
    },
    Person{
        Name: "proskuriakova",
        Description: "Tat???yana Avenirovna Proskuriakova (Russian: ???????????????? ?????????????????????? ??????????????????????????) (January 23 [O.S. January 10] 1909 ??? August 30, 1985) was a Russian-American Mayanist scholar and archaeologist who contributed significantly to the deciphering of Maya hieroglyphs, the writing system of the pre-Columbian Maya civilization of Mesoamerica.",
        Urls: []string{"https://en.wikipedia.org/wiki/Tatiana_Proskouriakoff",},
    },
    Person{
        Name: "ptolemy",
        Description: "Claudius Ptolemy. A Greco-Egyptian writer of Alexandria, known as a mathematician, astronomer, geographer, astrologer, and poet of a single epigram in the Greek Anthology.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ptolemy",},
    },
    Person{
        Name: "raman",
        Description: "C. V. Raman. Indian physicist who won the Nobel Prize in 1930 for proposing the Raman effect.",
        Urls: []string{"https://en.wikipedia.org/wiki/C._V._Raman",},
    },
    Person{
        Name: "ramanujan",
        Description: "Srinivasa Ramanujan. Indian mathematician and autodidact who made extraordinary contributions to mathematical analysis, number theory, infinite series, and continued fractions.",
        Urls: []string{"https://en.wikipedia.org/wiki/Srinivasa_Ramanujan",},
    },
    Person{
        Name: "ride",
        Description: "Sally Kristen Ride was an American physicist and astronaut. She was the first American woman in space, and the youngest American astronaut.",
        Urls: []string{"https://en.wikipedia.org/wiki/Sally_Ride",},
    },
    Person{
        Name: "ritchie",
        Description: "Dennis Ritchie, co-creator of UNIX and the C programming language.",
        Urls: []string{"https://en.wikipedia.org/wiki/Dennis_Ritchie",},
    },
    Person{
        Name: "rhodes",
        Description: "Ida Rhodes. American pioneer in computer programming, designed the first computer used for Social Security.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ida_Rhodes",},
    },
    Person{
        Name: "robinson",
        Description: "Julia Hall Bowman Robinson. American mathematician renowned for her contributions to the fields of computability theory and computational complexity theory.",
        Urls: []string{"https://en.wikipedia.org/wiki/Julia_Robinson",},
    },
    Person{
        Name: "roentgen",
        Description: "Wilhelm Conrad R??ntgen. German physicist who was awarded the first Nobel Prize in Physics in 1901 for the discovery of X-rays (R??ntgen rays).",
        Urls: []string{"https://en.wikipedia.org/wiki/Wilhelm_R%C3%B6ntgen",},
    },
    Person{
        Name: "rosalind",
        Description: "Rosalind Franklin. British biophysicist and X-ray crystallographer whose research was critical to the understanding of DNA.",
        Urls: []string{"https://en.wikipedia.org/wiki/Rosalind_Franklin",},
    },
    Person{
        Name: "rubin",
        Description: "Vera Rubin. American astronomer who pioneered work on galaxy rotation rates.",
        Urls: []string{"https://en.wikipedia.org/wiki/Vera_Rubin",},
    },
    Person{
        Name: "saha",
        Description: "Meghnad Saha. Indian astrophysicist best known for his development of the Saha equation, used to describe chemical and physical conditions in stars.",
        Urls: []string{"https://en.wikipedia.org/wiki/Meghnad_Saha",},
    },
    Person{
        Name: "sammet",
        Description: "Jean E. Sammet developed FORMAC, the first widely used computer language for symbolic manipulation of mathematical formulas.",
        Urls: []string{"https://en.wikipedia.org/wiki/Jean_E._Sammet",},
    },
    Person{
        Name: "sanderson",
        Description: "Mildred Sanderson. American mathematician best known for Sanderson's theorem concerning modular invariants.",
        Urls: []string{"https://en.wikipedia.org/wiki/Mildred_Sanderson",},
    },
    Person{
        Name: "satoshi",
        Description: "Satoshi Nakamoto is the name used by the unknown person or group of people who developed bitcoin, authored the bitcoin white paper, and created and deployed bitcoin's original reference implementation.",
        Urls: []string{"https://en.wikipedia.org/wiki/Satoshi_Nakamoto",},
    },
    Person{
        Name: "shamir",
        Description: "Adi Shamir. Israeli cryptographer whose numerous inventions and contributions to cryptography include the Ferge Fiat Shamir identification scheme, the Rivest Shamir Adleman (RSA) public-key cryptosystem, the Shamir's secret sharing scheme, the breaking of the Merkle-Hellman cryptosystem, the TWINKLE and TWIRL factoring devices and the discovery of differential cryptanalysis (with Eli Biham).",
        Urls: []string{"https://en.wikipedia.org/wiki/Adi_Shamir",},
    },
    Person{
        Name: "shannon",
        Description: "Claude Shannon. The father of information theory and founder of digital circuit design theory.",
        Urls: []string{"https://en.wikipedia.org/wiki/Claude_Shannon)",},
    },
    Person{
        Name: "shaw",
        Description: "Carol Shaw. Originally an Atari employee, Carol Shaw is said to be the first female video game designer.",
        Urls: []string{"https://en.wikipedia.org/wiki/Carol_Shaw_(video_game_designer)",},
    },
    Person{
        Name: "shirley",
        Description: "Dame Stephanie \"Steve\" Shirley. Founded a software company in 1962 employing women working from home.",
        Urls: []string{"https://en.wikipedia.org/wiki/Steve_Shirley",},
    },
    Person{
        Name: "shockley",
        Description: "William Shockley co-invented the transistor.",
        Urls: []string{"https://en.wikipedia.org/wiki/William_Shockley",},
    },
    Person{
        Name: "shtern",
        Description: "Lina Solomonovna Stern (or Shtern; Russian: ???????? ?????????????????????? ??????????; 26 August 1878 ??? 7 March 1968) was a Soviet biochemist, physiologist and humanist whose medical discoveries saved thousands of lives at the fronts of World War II. She is best known for her pioneering work on blood???brain barrier, which she described as hemato-encephalic barrier in 1921.",
        Urls: []string{"https://en.wikipedia.org/wiki/Lina_Stern",},
    },
    Person{
        Name: "sinoussi",
        Description: "Fran??oise Barr??-Sinoussi. French virologist and Nobel Prize Laureate in Physiology or Medicine; her work was fundamental in identifying HIV as the cause of AIDS.",
        Urls: []string{"https://en.wikipedia.org/wiki/Fran%C3%A7oise_Barr%C3%A9-Sinoussi",},
    },
    Person{
        Name: "snyder",
        Description: "Betty Snyder. One of the original programmers of the ENIAC.",
        Urls: []string{"https://en.wikipedia.org/wiki/ENIAC", "https://en.wikipedia.org/wiki/Betty_Holberton",},
    },
    Person{
        Name: "solomon",
        Description: "Cynthia Solomon. Pioneer in the fields of artificial intelligence, computer science and educational computing. Known for creation of Logo, an educational programming language.",
        Urls: []string{"https://en.wikipedia.org/wiki/Cynthia_Solomon",},
    },
    Person{
        Name: "spence",
        Description: "Frances Spence. One of the original programmers of the ENIAC.",
        Urls: []string{"https://en.wikipedia.org/wiki/ENIAC", "https://en.wikipedia.org/wiki/Frances_Spence",},
    },
    Person{
        Name: "stonebraker",
        Description: "Michael Stonebraker is a database research pioneer and architect of Ingres, Postgres, VoltDB and SciDB. Winner of 2014 ACM Turing Award.",
        Urls: []string{"https://en.wikipedia.org/wiki/Michael_Stonebraker",},
    },
    Person{
        Name: "sutherland",
        Description: "Ivan Edward Sutherland. American computer scientist and Internet pioneer, widely regarded as the father of computer graphics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ivan_Sutherland",},
    },
    Person{
        Name: "swanson",
        Description: "Janese Swanson (with others) developed the first of the Carmen Sandiego games. She went on to found Girl Tech.",
        Urls: []string{"https://en.wikipedia.org/wiki/Janese_Swanson",},
    },
    Person{
        Name: "swartz",
        Description: "Aaron Swartz was influential in creating RSS, Markdown, Creative Commons, Reddit, and much of the internet as we know it today. He was devoted to freedom of information on the web.",
        Urls: []string{"https://en.wikiquote.org/wiki/Aaron_Swartz",},
    },
    Person{
        Name: "swirles",
        Description: "Bertha Swirles was a theoretical physicist who made a number of contributions to early quantum theory.",
        Urls: []string{"https://en.wikipedia.org/wiki/Bertha_Swirles",},
    },
    Person{
        Name: "taussig",
        Description: "Helen Brooke Taussig. American cardiologist and founder of the field of paediatric cardiology.",
        Urls: []string{"https://en.wikipedia.org/wiki/Helen_B._Taussig",},
    },
    Person{
        Name: "tereshkova",
        Description: "Valentina Tereshkova is a Russian engineer, cosmonaut and politician. She was the first woman to fly to space in 1963. In 2013, at the age of 76, she offered to go on a one-way mission to Mars.",
        Urls: []string{"https://en.wikipedia.org/wiki/Valentina_Tereshkova",},
    },
    Person{
        Name: "tesla",
        Description: "Nikola Tesla invented the AC electric system and every gadget ever used by a James Bond villain.",
        Urls: []string{"https://en.wikipedia.org/wiki/Nikola_Tesla",},
    },
    Person{
        Name: "tharp",
        Description: "Marie Tharp. American geologist and oceanic cartographer who co-created the first scientific map of the Atlantic Ocean floor. Her work led to the acceptance of the theories of plate tectonics and continental drift.",
        Urls: []string{"https://en.wikipedia.org/wiki/Marie_Tharp",},
    },
    Person{
        Name: "thompson",
        Description: "Ken Thompson, co-creator of UNIX and the C programming language.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ken_Thompson",},
    },
    Person{
        Name: "torvalds",
        Description: "Linus Torvalds invented Linux and Git.",
        Urls: []string{"https://en.wikipedia.org/wiki/Linus_Torvalds",},
    },
    Person{
        Name: "tu",
        Description: "Youyou Tu. Chinese pharmaceutical chemist and educator known for discovering artemisinin and dihydroartemisinin, used to treat malaria, which has saved millions of lives. Joint winner of the 2015 Nobel Prize in Physiology or Medicine.",
        Urls: []string{"https://en.wikipedia.org/wiki/Tu_Youyou",},
    },
    Person{
        Name: "turing",
        Description: "Alan Turing was a founding father of computer science.",
        Urls: []string{"https://en.wikipedia.org/wiki/Alan_Turing.",},
    },
    Person{
        Name: "varahamihira",
        Description: "Varahamihira. Ancient Indian mathematician who discovered trigonometric formulae during 505-587 CE.",
        Urls: []string{"https://en.wikipedia.org/wiki/Var%C4%81hamihira#Contributions",},
    },
    Person{
        Name: "vaughan",
        Description: "Dorothy Vaughan was a NASA mathematician and computer programmer on the SCOUT launch vehicle program that put America's first satellites into space.",
        Urls: []string{"https://en.wikipedia.org/wiki/Dorothy_Vaughan",},
    },
    Person{
        Name: "villani",
        Description: "C??dric Villani. French mathematician, won Fields Medal, Fermat Prize and Poincar?? Price for his work in differential geometry and statistical mechanics.",
        Urls: []string{"https://en.wikipedia.org/wiki/C%C3%A9dric_Villani",},
    },
    Person{
        Name: "visvesvaraya",
        Description: "Sir Mokshagundam Visvesvaraya, is a notable Indian engineer. He is a recipient of the Indian Republic's highest honour, the Bharat Ratna, in 1955. On his birthday, 15 September is celebrated as Engineer's Day in India in his memory.",
        Urls: []string{"https://en.wikipedia.org/wiki/Visvesvaraya",},
    },
    Person{
        Name: "volhard",
        Description: "Christiane N??sslein-Volhard. German biologist, won Nobel Prize in Physiology or Medicine in 1995 for research on the genetic control of embryonic development.",
        Urls: []string{"https://en.wikipedia.org/wiki/Christiane_N%C3%BCsslein-Volhard",},
    },
    Person{
        Name: "wescoff",
        Description: "Marlyn Wescoff. One of the original programmers of the ENIAC.",
        Urls: []string{"https://en.wikipedia.org/wiki/ENIAC", "https://en.wikipedia.org/wiki/Marlyn_Meltzer",},
    },
    Person{
        Name: "wilbur",
        Description: "Sylvia B. Wilbur. British computer scientist who helped develop the ARPANET, was one of the first to exchange email in the UK and a leading researcher in computer-supported collaborative work.",
        Urls: []string{"https://en.wikipedia.org/wiki/Sylvia_Wilbur",},
    },
    Person{
        Name: "wiles",
        Description: "Andrew Wiles. Notable British mathematician who proved the enigmatic Fermat's Last Theorem.",
        Urls: []string{"https://en.wikipedia.org/wiki/Andrew_Wiles",},
    },
    Person{
        Name: "williams",
        Description: "Roberta Williams, did pioneering work in graphical adventure games for personal computers, particularly the King's Quest series.",
        Urls: []string{"https://en.wikipedia.org/wiki/Roberta_Williams",},
    },
    Person{
        Name: "williamson",
        Description: "Malcolm John Williamson. British mathematician and cryptographer employed by the GCHQ. Developed in 1974 what is now known as Diffie-Hellman key exchange (Diffie and Hellman first published the scheme in 1976).",
        Urls: []string{"https://en.wikipedia.org/wiki/Malcolm_J._Williamson",},
    },
    Person{
        Name: "wilson",
        Description: "Sophie Wilson designed the first Acorn Micro-Computer and the instruction set for ARM processors.",
        Urls: []string{"https://en.wikipedia.org/wiki/Sophie_Wilson",},
    },
    Person{
        Name: "wing",
        Description: "Jeannette Wing, co-developed the Liskov substitution principle.",
        Urls: []string{"https://en.wikipedia.org/wiki/Jeannette_Wing",},
    },
    Person{
        Name: "wozniak",
        Description: "Steve Wozniak invented the Apple I and Apple II.",
        Urls: []string{"https://en.wikipedia.org/wiki/Steve_Wozniak",},
    },
    Person{
        Name: "wright",
        Description: "The Wright brothers, Orville and Wilbur, credited with inventing and building the world's first successful airplane and making the first controlled, powered and sustained heavier-than-air human flight.",
        Urls: []string{"https://en.wikipedia.org/wiki/Wright_brothers",},
    },
    Person{
        Name: "wu",
        Description: "Chien-Shiung Wu. Chinese-American experimental physicist who made significant contributions to nuclear physics.",
        Urls: []string{"https://en.wikipedia.org/wiki/Chien-Shiung_Wu",},
    },
    Person{
        Name: "yalow",
        Description: "Rosalyn Sussman Yalow. Rosalyn Sussman Yalow was an American medical physicist, and a co-winner of the 1977 Nobel Prize in Physiology or Medicine for development of the radioimmunoassay technique.",
        Urls: []string{"https://en.wikipedia.org/wiki/Rosalyn_Sussman_Yalow",},
    },
    Person{
        Name: "yonath",
        Description: "Ada Yonath. An Israeli crystallographer, the first woman from the Middle East to win a Nobel prize in the sciences.",
        Urls: []string{"https://en.wikipedia.org/wiki/Ada_Yonath",},
    },
    Person{
        Name: "zhukovsky",
        Description: "Nikolay Yegorovich Zhukovsky (Russian: ???????????????? ?????????????????? ????????????????????, January 17 1847 ??? March 17, 1921) was a Russian scientist, mathematician and engineer, and a founding father of modern aero- and hydrodynamics. Whereas contemporary scientists scoffed at the idea of human flight, Zhukovsky was the first to undertake the study of airflow. He is often called the Father of Russian Aviation.",
        Urls: []string{"https://en.wikipedia.org/wiki/Nikolay_Yegorovich_Zhukovsky",},
    },
}

// picks a random adjective and person
func GetRandomName() Person {
begin:
    adjective := adjectives[rand.Intn(len(adjectives))]
    person := persons[rand.Intn(len(persons))]

    person.Adjective = adjective

    // Steve Wozniak is not boring
    if fmt.Sprintf("%s%s", person.Adjective, person.Name) == "boringwozniak" {
        goto begin
    }

    return person
}

