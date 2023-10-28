INSERT INTO client (id, "name", cpf, email)
VALUES (DEFAULT,'Lavínia', '99283945042', 'lavinia_mirella_monteiro@teadit.com.br'),
       (DEFAULT,'Benjamin', '30631030638', 'benjaminfarias@negocios-de-valor.com'),
       (DEFAULT,'Rafaela Carolina', '39512179873', 'rafaela.carolina.darocha@kimmay.com.br'),
       (DEFAULT,'Cecília Eduarda', '11047796503', 'cecilia_assis@bseletronicos.com.br'),
       (DEFAULT,'Antonio', '90155027573', 'antonio-dossantos73@djapan.com.br'),
       (DEFAULT,'Carlos Isaac', '90383229332', 'carlos_rezende@oana.com.br'),
       (DEFAULT,'Luana Rocha', '06609642877', 'luana.luciana.rocha@knowconsulting.com.br'),
       (DEFAULT,'Gabriela', '36012078773', NULL),
       (DEFAULT,'Caio', NULL, 'caio-carvalho88@aulicinobastos.com.br'),
       (DEFAULT,'Henrique', NULL, NULL);

INSERT INTO product (id, "name", description, category, image, price)
VALUES (DEFAULT,'Big Mac', 'Dois hambúrgueres (100% carne bovina), alface americana, queijo sabor cheddar, molho especial, cebola, picles e pão com gergelim.', 'LANCHE', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kzXCTbnv/200/200/original?country=br', 18.00),
        (DEFAULT,'Duplo Cheddar McMelt', 'Dois hambúrgueres (100% carne bovina), molho lácteo cremoso sabor cheddar, cebola ao molho shoyu e pão escuro com gergelim.', 'LANCHE', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kzXWKJ6A/200/200/original?country=br', 20.00),
        (DEFAULT,'Duplo Burger Bacon', 'Dois hambúrgueres (100% carne bovina), queijo sabor cheddar, cebola, fatias de bacon, ketchup, mostarda e pão com gergelim.', 'LANCHE', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kMX5kx4H/200/200/original?country=br', 18.00),
        (DEFAULT,'McFritas Grande', 'A batata frita mais famosa do mundo. Deliciosas batatas selecionadas, fritas, crocantes por fora, macias por dentro, douradas, irresistíveis, saborosas, famosas, e todos os outros adjetivos positivos que você quiser dar.', 'ACOMPANHAMENTO', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kUXVg4F7/200/200/original?country=br', 10.00),
        (DEFAULT,'McFritas Média', 'A batata frita mais famosa do mundo. Deliciosas batatas selecionadas, fritas, crocantes por fora, macias por dentro, douradas, irresistíveis, saborosas, famosas, e todos os outros adjetivos positivos que você quiser dar.', 'ACOMPANHAMENTO', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kUXGZHtB/200/200/original?country=br', 8.00),
        (DEFAULT,'McFritas Pequena', 'A batata frita mais famosa do mundo. Deliciosas batatas selecionadas, fritas, crocantes por fora, macias por dentro, douradas, irresistíveis, saborosas, famosas, e todos os outros adjetivos positivos que você quiser dar.', 'ACOMPANHAMENTO', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kUXgPmuC/200/200/original?country=br', 7.00),
        (DEFAULT,'Coca-Cola 300ml', 'Refrescante e geladinha. Uma bebida assim refresca a vida. Você pode escolher entre Coca-Cola, Coca-Cola Zero, Sprite sem Açúcar, Fanta Guaraná e Fanta Laranja.', 'BEBIDA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kNXZJR6V/200/200/original?country=br', 12.00),
        (DEFAULT,'Coca-Cola 500ml', 'Refrescante e geladinha. Uma bebida assim refresca a vida. Você pode escolher entre Coca-Cola, Coca-Cola Zero, Sprite sem Açúcar, Fanta Guaraná e Fanta Laranja.', 'BEBIDA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kNXBvqQj/200/200/original?country=br', 16.00),
        (DEFAULT,'Coca-Cola 700ml', 'Refrescante e geladinha. Uma bebida assim refresca a vida. Você pode escolher entre Coca-Cola, Coca-Cola Zero, Sprite sem Açúcar, Fanta Guaraná e Fanta Laranja.', 'BEBIDA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kNXMLd8s/200/200/original?country=br', 20.00),
        (DEFAULT,'Água Mineral', 'Água sem gás.', 'BEBIDA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$k7X5DQ6J/200/200/original?country=br', 6.00),
        (DEFAULT,'Del Valle Laranja 500ml', 'Deliciosos sabores à sua escolha. Néctar de fruta nos sabores uva ou laranja.', 'BEBIDA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kNXkpLzq/200/200/original?country=br', 15.00),
        (DEFAULT,'Del Valle Laranja 700ml', 'Deliciosos sabores à sua escolha. Néctar de fruta nos sabores uva ou laranja.', 'BEBIDA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kNXWVFLM/200/200/original?country=br', 19.00),
        (DEFAULT,'Torta de Maçã', 'Boa demais. Parece a receita lá de casa. Massa quentinha e crocante envolvendo deliciosos recheios de banana ou maçã com gostinho de doce caseiro.', 'SOBREMESA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$krXTZ9Ue/200/200/original?country=br', 5.00),
        (DEFAULT,'McShake Ovomaltine', 'Deliciosamente cremoso. O novo McShake Ovomaltine é feito com leite e batido na hora. Uma delícia!', 'SOBREMESA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kJX0TX33/200/200/original?country=br', 12.00),
        (DEFAULT,'Casquinha Chocolate', 'A sobremesa que o Brasil todo adora. Uma casquinha supercrocante, com bebida láctea sabor chocolate que vai bem a qualquer hora.', 'SOBREMESA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kpXyfJ7k/200/200/original?country=br', 4.00),
        (DEFAULT,'Casquinha Baunilha', 'A sobremesa que o Brasil todo adora. Uma casquinha supercrocante, com bebida láctea sabor chocolate que vai bem a qualquer hora.', 'SOBREMESA', 'https://cache-backend-mcd.mcdonaldscupones.com/media/image/product$kpXnFFzy/200/200/original?country=br', 4.00);

INSERT INTO voucher (id, code, percentage, expires_at)
VALUES (DEFAULT,'BEMVINDO10', 10.00, '2023-11-30 23:59:59'),
        (DEFAULT,'OFERTAS50', 50.00, '2023-11-01 23:59:59'),
        (DEFAULT,'OGERENTEFICOULOUCO', 100.00, '2023-10-21 23:59:59');

INSERT INTO orders (id, status, verification_code, client_id, voucher_id)
VALUES (DEFAULT,'RECEBIDO', 'abc001', 1, NULL),
       (DEFAULT,'EM PREPARAÇÃO', 'bcd002', 2, 1),
       (DEFAULT,'PRONTO', 'cde003', 3, 2),
       (DEFAULT,'FINALIZADO', 'def004', 3, NULL);

INSERT INTO orders_products (id, orders_id, product_id, quantity, total_price, discount)
VALUES (DEFAULT,1, 1, 2, 36.00, 0.00),
       (DEFAULT,1, 8, 1, 16.00, 0.00),
       (DEFAULT,1, 14, 1, 5.00, 0.00),
       (DEFAULT,2, 2, 1, 20.00, 2.00),
       (DEFAULT,2, 10, 1, 6.00, 0.60),
       (DEFAULT,3, 15, 1, 4.00, 2.00),
       (DEFAULT,4, 3, 1, 18.00, 0.00),
       (DEFAULT,4, 12, 1, 19.00, 0.00);

