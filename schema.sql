CREATE TABLE blogs (
    id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE entries (
    id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
    blog_id SMALLINT UNSIGNED NOT NULL REFERENCES blogs(id),
    author VARCHAR(50) DEFAULT "Anonymous",
    title VARCHAR(100) NOT NULL DEFAULT "Untitled Entry",
    content TEXT,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    edited TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO blogs 
SET
        title = "first blog";

INSERT INTO blogs 
SET
        title = "second blog";

INSERT INTO blogs 
SET
        title = "third blog";

INSERT INTO blogs 
SET
        title = "fourth blog";

INSERT INTO entries (author, content, title, blog_id)
VALUES ("Autho Rone", "This is the first blog post I've done
        It's a really good one! I think it's super neat to be
        blogging in this weird format!", "My New Blog!!", 1);

INSERT INTO entries (author, content, title, blog_id)
VALUES ("Thirdaut Hor", "There! I have posted a blog! I hope you are happy!", "Vindictive Post", 1);

INSERT INTO entries (author, content, title, blog_id)
VALUES ("Secon Dauthor", "Neutra polaroid next level, Etsy helvetica hoodie ennui fanny pack. Quinoa thundercats PBR retro. Roof party swag iPhone tousled fingerstache skateboard direct trade, flannel gluten-free. Readymade retro High Life sartorial Bushwick scenester. Cred lo-fi beard locavore stumptown. Fap mustache fanny pack, Portland retro banjo butcher High Life Truffaut sartorial scenester. Organic farm-to-table ennui stumptown retro, pork belly shoreditch butcher.

Cliche Terry Richardson tofu meh pour-over biodiesel, ethical thundercats. YOLO gastropub chambray lo-fi, tousled readymade vinyl dreamcatcher pork belly. Mlkshk stumptown Portland fap, squid four loko small batch biodiesel blog. Banh mi Terry Richardson vinyl, occupy master cleanse intelligentsia hella Echo Park aesthetic craft beer synth squid narwhal +1 literally. Bushwick keffiyeh trust fund kitsch Carles ethnic. Gastropub banh mi keytar, church-key artisan actually tumblr. Put a bird on it pug synth, Banksy raw denim meggings YOLO Etsy lomo banjo american apparel four loko aesthetic.

Aesthetic american apparel cardigan, DIY photo booth 90's cred blue bottle small batch mumblecore organic scenester artisan kale chips. Roof party bicycle rights asymmetrical, locavore freegan try-hard semiotics sustainable Bushwick typewriter Godard tofu. Occupy banjo trust fund Neutra synth. Vinyl jean shorts deep v art party, Wes Anderson Williamsburg flannel narwhal ugh before they sold out letterpress semiotics. Retro Wes Anderson farm-to-table, hoodie craft beer bicycle rights Truffaut mlkshk hella umami put a bird on it jean shorts cliche. Umami Austin master cleanse, Tonx tattooed vegan fap small batch gastropub Vice typewriter. Marfa quinoa pug, you probably haven't heard of them vinyl ugh hella tote bag post-ironic flexitarian iPhone skateboard.

Austin Echo Park umami, skateboard freegan YOLO cray squid. Mixtape hoodie scenester tote bag meh literally. Squid Marfa ennui, Schlitz pug direct trade Carles sriracha biodiesel selfies single-origin coffee shoreditch pitchfork gentrify. You probably haven't heard of them ugh Truffaut YOLO pitchfork disrupt. Try-hard trust fund intelligentsia ethical, gastropub pug raw denim butcher shoreditch. Blue bottle kale chips chambray Odd Future beard, Bushwick trust fund hella. Master cleanse Pinterest Portland twee, fashion axe pop-up pitchfork selfies irony 90's direct trade Tonx intelligentsia.

", "My Hipster Blog Post", 2);
