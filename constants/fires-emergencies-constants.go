package constants

func GetFireMapCodes() map[string]string {
	fireCodesMap := map[string]string{

		"1-1":  "INCIDENTE DE INCENDIO ESTRUCTURAL SIMPLE",
		"1-2":  "INCIDENTE DE INCENDIO ESTRUCTURAL COMPLEJO",
		"1-3":  "INCIDENTE DE INCENDIO ESTRUCTURAL EN EDIFICIO DE ALTURA",
		"1-4":  "INCIDENTE DE INCENDIO ESTRUCTURAL CON AFLUENCIA DE PÚBLICO",
		"2-1":  "INCIDENTE DE INCENDIO FORESTAL DE INTERFASE",
		"2-2":  "INCIDENTE DE INCENDIO FORESTAL",
		"2-3":  "INCIDENTE DE INCENDIO EN ESCOMBROS O BASURA",
		"3-1":  "INCIDENTE DE INCENDIO VEHICULAR MENOR",
		"3-2":  "INCIDENTE DE INCENDIO VEHICULAR MAYOR",
		"3-3":  "INCIDENTE DE INCENDIO VEHICULAR COMPLEJO",
		"4-1":  "INCIDENTE CON GASES COMBUSTIBLES SIMPLE",
		"4-2":  "INCIDENTE CON GASES COMBUSTIBLES INTERMEDIO",
		"4-3":  "INCIDENTE CON GASES COMBUSTIBLES COMPLEJOS",
		"4-4":  "INCIDENTE CON MATERIALES PELIGROSOS",
		"5-1":  "INCIDENTE DE RESCATE VEHICULAR SIMPLE",
		"5-2":  "INCIDENTE DE RESCATE VEHICULAR COMPLEJO",
		"5-3":  "INCIDENTE DE RESCATE VEHICULAR EN TÚNEL",
		"6-1":  "INCIDENTE DE RESCATE",
		"6-2":  "INCIDENTE DE RESCATE EN DESNIVEL",
		"6-3":  "INCIDENTE DE RESCATE EN ALTURA",
		"6-4":  "INCIDENTE DE RESCATE EN ESPACIOS CONFINADOS",
		"6-5":  "INCIDENTE DE RESCATE URBANO",
		"6-6":  "INCIDENTE DE RESCATE AGRESTE",
		"7":    "ACUARTELAMIENTO GENERAL",
		"7-1":  "ACUARTELAMIENTO A COMPAÑÍAS DE ESPECIALIDAD",
		"7-2":  "ACUARTELAMIENTO A GRUPO Y/O EQUIPO ESPECIALIZADO",
		"8":    "APOYO A OTROS CUERPOS DE BOMBEROS",
		"9-1":  "INCIDENTE DE INCENDIO INDUSTRIAL",
		"9-2":  "INCIDENTE CON MULTIPLICIDAD DE PERSONAS LESIONADAS",
		"9-3":  "INCIDENTE AÉREO",
		"9-4":  "INCIDENTE FERROVIARIO",
		"9-5":  "INCIDENTE POR INCENDIO EN VERTEDERO O RELLENO SANITARIO",
		"10":   "INCIDENTE ELÉCTRICO",
		"11":   "INCIDENTE NO CATEGORIZADO",
		"12":   "OTROS SERVICIOS",
		"13":   "REBROTE DE INCIDENTE",
		"14":   "SIMULACRO DE INCIDENTE",
		"15":   "CITACIÓN INSTITUCIONAL",
		"10-0": "Llamado estructural",

		"10-0-1":   "FUEGO EN FASE INICIAL, DE INCREMENTO O LIBRE EN EL SIGUIENTE TIPO DE ESTRUCTURA",
		"10-0-1-1": "CASA HABITACIÓN 1 PISO (INCLUYE KIOSKOS, GARITAS, CASETAS)",
		"10-0-1-2": "CASA HABITACIÓN MÁS DE 1 PISO y HASTA 4 PISOS",
		"10-0-1-3": "CASA HABITACIÓN INFORMAL y/o DE MATERIAL DE ALTA COMBUSTIBILIDAD (MEJORAS, MEDIAGUAS, RUCOS Y SIMILARES)",
		"10-0-2":   "FUEGO EN BASE INICIAL EN EDIFICACIONES EN ALTURA",
		"10-0-2-1": "EDIFICIOS HABITACIONALES O DE OFICINAS ENTRE EL 1º AL 5º PISO",
		"10-0-2-2": "EDIFICIOS HABITACIONALES O DE OFICINAS ENTRE EL 6º A 10º PISO",
		"10-0-2-3": "EDIFICIOS HABITACIONALES O DE OFICINAS ENTRE EL 11º A 15º PISO",
		"10-0-2-4": "EDIFICIOS HABITACIONALES O DE OFICINAS ENTRE EL 16º A 20º PISO",
		"10-0-2-5": "EDIFICIOS HABITACIONALES O DE OFICINAS DESDE EL PISO 21º HACIA ARRIBA",
		"10-0-2-6": "SUBTERRÁNEOS DE EDIFICIOS HABITACIONALES O DE OFICINAS DE CUALQUIER ALTURA",
		"10-0-3":   "LUGARES CON ALTA AFLUENCIA DE PÚBLICO",
		"10-0-3-1": "LOCAL COMERCIAL DE UN PISO",
		"10-0-3-2": "SUPERMERCADO",
		"10-0-3-3": "CENTROS COMERCIALES (MALL, FERIAS LIBRES, MERCADOS)",
		"10-0-3-4": "LUGAR DE ALOJAMIENTO COMERCIAL (HOTEL, MOTEL, RESIDENCIALES, HOSTALES)",
		"10-0-3-5": "ESTABLECIMIENTOS EDUCACIONALES",
		"10-0-3-6": "EDIFICIOS PÚBLICOS (MUNICIPALIDAD, INTENDENCIA, GOBERNACIÓN, SERVICIOS, MINISTERIOS, ENTRE OTROS)",
		"10-0-3-7": "TEMPLOS / IGLESIAS, LUGARES DE CULTO",
		"10-0-3-8": "TERMINAL RODOVIARIO (BUSES, FERROCARRILES)",
		"10-0-3-9": "RECINTOS ONG (CRUZ ROJA, DEFENSA CIVIL, BOMBEROS DE CHILE)",
		"10-0-4":   "LUGARES CON PERSONAS CON MOVILIDAD RESTRINGIDA",
		"10-0-4-1": "HOSPITAL O CLÍNICAS",
		"10-0-4-2": "CONSULTORIOS / POLICLÍNICOS",

		"10-0-4-3": "CÁRCEL, CENTRO DE REHABILITACION O RECINTOS DE GENDARMERIA/ CENTROS PRIVATIVOS DE LIBERTAD",
		"10-0-4-4": "HOGARES DE ACOGIDA E INTERNADOS",

		"10-0-5":    "LUGARES CON RIESGOS DE INCENDIOS ESPECIALES Y/O DE ALTA CARGA COMBUSTIBLE",
		"10-0-5-1":  "INDUSTRIAS",
		"10-0-5-2":  "PLANTAS FAENADORAS (ALIMENTOS, FRUTÍCOLAS)",
		"10-0-5-3":  "POLVORINES",
		"10-0-5-4":  "BODEGA",
		"10-0-5-5":  "TALLERES ARTESANALES",
		"10-0-5-6":  "BARRACAS O ASERRADEROS",
		"10-0-5-7":  "CENTROS DE ACOPIOS (LUGAR ABIERTO DE ALMACENAMIENTO)",
		"10-0-5-8":  "CENTRAL GENERADORA O SUBESTACIÓN ELÉCTRICAS",
		"10-0-5-9":  "SERVICENTROS",
		"10-0-5-10": "RECINTO PORTUARIOS",
		"10-0-5-11": "FUEGO EN MINAS",
		"10-0-5-12": "RECINTO AERONÁUTICO",
		"10-0-5-13": "CENTROS DE REACREACIÓN MASIVA",
		"10-0-5-14": "RECINTOS DEPORTIVOS",
		"10-0-5-15": "RECINTOS DE FF.AA. POLICIALES",
		"10-0-5-16": "SUBTERRÁNEOS",
		"10-0-5-17": "CÁMARAS SUBTERRÁNEAS",
		"10-0-5-18": "TÚNELES VIALES",

		"10-0-5-19": "ALUMBRADO PÚBLICO, ACOMETIDAS, TRANSFORMADORES ELÉCTRICOS O SIMILARES",
		"10-0-5-20": "CORRALES ESTABLOS O CANILES",

		"10-1":    "FUEGO EN FASE INICIAL DE INCREMENTO O LIBRE, EN VEHÍCULOS Y MEDIOS DE TRANSPORTE",
		"10-1-1":  "VEHÍCULOS, AUTOS, CAMIONETAS, MOTOCICLETAS",
		"10-1-2":  "VEHÍCULOS CAMIONES (VEHÍCULOS MAYORES A 3.500 KG.)",
		"10-1-3":  "VEHÍCULOS BUSES URBANOS (MICROBUSES, TRANSPORTE ESCOLAR)",
		"10-1-4":  "VEHÍCULOS BUSES INTERURBANOS",
		"10-1-5":  "VEHÍCULOS MILITAR ESPECIAL",
		"10-1-6":  "MAQUINARIA AGRICOLA / PESADA/ MINERA",
		"10-1-7":  "VEHÍCULOS TRACCIÓN ANIMAL O HUMANA",
		"10-1-8":  "EMBARCACIONES MENORES HASTA 50 TON.",
		"10-1-9":  "EMBARCACIONES MAYORES SOBRE 50 TON.",
		"10-1-10": "AERONAVES (AVIONETAS, HELICÓPTEROS, AVIONES)",
		"10-1-11": "FERROCARRILES (LOCOMOTORA, CARROS)",
		"10-1-12": "TREN SUBTERRÁNEO Y/O URBANO",

		"10-2":    "FUEGO EN ÁREAS ABIERTAS O NO CONFINADAS, PÚBLICAS O PRIVADAS, DE LOS SIGUIENTES TIPOS",
		"10-2-1":  "PASTIZALES URBANOS",
		"10-2-2":  "PASTIZALES RURALES",
		"10-2-3":  "PLANTACIÓN AGRÍCOLA",
		"10-2-4":  "PÁRQUES NACIONALES",
		"10-2-5":  "INTERFASE (BOSQUE POBLACIÓN)",
		"10-2-6":  "MICROBASURALES",
		"10-2-7":  "DESECHOS, ESCOMBROS",
		"10-2-8":  "BOSQUES (ÁRBOLES, MATORRALES)",
		"10-2-9":  "VERTEDERO O RELLENOS SANITARIOS",
		"10-2-10": "BIEN DE USO PÚBLICO (PARADEROS, ORNAMENTACIÓN, LUMINARIAS, OTROS)",
		"10-3":    "RESCATE DE VÍCTIMAS",
		"10-3-1":  "RESCATE DESDE ASCENSORES, ESCALAS MECÁNICAS",
		"10-3-2":  "RESCATE DESDE ESTRUCTURAS COLAPSADAS",
		"10-3-3":  "RESCATE DESDE DESNIVEL URBANO (MÁS DE 2 PISOS HACIA ARRIBA)",
		"10-3-4":  "RESCATE DESDE DESNIVEL BAJO (ALCANTARILLA, CÁMARAS, POZOS, ETC.)",
		"10-3-5":  "RESCATE DESDE DESNIVEL RURAL (QUEBRADA, ACANTILADO O SIMILAR)",
		"10-3-6":  "RESCATE DESDE MINAS SUBTERRÁNEAS",
		"10-3-7":  "RESCATE DE AGUAS EN MOVIMIENTO (RÍOS, CANALES, ESTEROS)",
		"10-3-8":  "RESCATE DESDE AGUAS QUIETAS (LAGOS, LAGUNAS, SIMILARES, PISCINAS)",
		"10-3-9":  "RESCATE DESDE EL MAR",

		"10-3-10": "RESCATE DESDE MAQUINARIAS (PARTE DE UNA PERSONA ATRAPADA EN MÁQUINAS O SIMILARES)",
		"10-3-11": "RESCATE EN VÍA PÚBLICA POR ATROPELLO",
		"10-3-12": "RESCATE POR ALUVIONES Y DESLIZAMIENTOS",
		"10-3-13": "RESCATE EN VOLCÁN Y/O MONTAÑA SOBRE 700 MTS.",
		"10-3-14": "RESCATE EN CERRO",
		"10-3-15": "RESCATE EN NIEVE O GLACIAR",
		"10-3-16": "RESCATE DE VEHÍCULOS EN LUGARES RIESGOSOS",
		"10-3-17": "RESCATE EN AERONAVES CAÍDAS O FUERA DE RECINTOS AERONÁUTICOS",
		"10-3-18": "RESCATE DE PERSONAS ATRAPADAS O EMPALADAS",
		"10-3-19": "RESCATE POR ELECTROCUCIÓN EN LUGARES PÚBLICOS O PRIVADOS",

		"10-4":    "RESCATE DE VÍCTIMAS DESDE LOS SIGUIENTES TIPOS DE VEHÍCULOS Y EN LAS SIGUIENTES CONDICIONES: (VOLCAMIENTO O DESPISTE SE CONSIDERA CHOQUE)",
		"10-4-1":  "CHOQUE DE AUTO",
		"10-4-2":  "CHOQUE DE CAMIÓN",
		"10-4-3":  "CHOQUE DE BUS",
		"10-4-4":  "CHOQUE DE MOTOCICLETAS, TRACCIÓN HUMANA O ANIMAL",
		"10-4-5":  "CHOQUE DE MAQUINARIAS AGRÍCOLAS, PESADAS Y MINERAS",
		"10-4-6":  "CHOQUE FERROVIARIO",
		"10-4-7":  "CHOQUE DE VEHÍCULOS DE EMERGENCIA",
		"10-4-8":  "CHOQUE DE VEHÍCULOS DE FF.AA. Y POLICIALES",
		"10-4-9":  "CHOQUE DE EMBARCACIONES",
		"10-4-10": "CHOQUE DE TRENES SUBTERRÁNEOS",
		"10-4-11": "CHOQUE DE VEHÍCULOS PESADOS, AGRÍCOLAS, MINEROS",
		"10-4-12": "COLISIÓN AUTOS",
		"10-4-13": "COLISIÓN CAMIONES",
		"10-4-14": "COLISIÓN DE BUSES",
		"10-4-15": "COLISIÓN DE MOTOCICLETAS, TRACCIÓN HUMANA O ANIMAL",
		"10-4-16": "COLISIÓN DE MAQUINARIAS AGRÍCOLAS, PESADAS Y MINERAS",
		"10-4-17": "COLISIÓN FERROVIARIA",
		"10-4-18": "COLISIÓN DE VEHÍCULOS DE EMERGENCIA",
		"10-4-19": "COLISIÓN DE VEHÍCULOS DE FF.AA. Y POLICIALES",
		"10-4-20": "COLISIÓN DE EMBARCACIONES",
		"10-4-21": "COLISIÓN TRENES SUBTERRÁNEOS",
		"10-4-22": "COLISION DE VEHÍCULOS PESADOS, AGRÍCOLAS, MINEROS",
		"10-4-23": "DESBARRANCAMIENTO DE VEHÍCULOS",

		"10-5":   "ACTUACIÓN EN ACCIDENTES CON LAS SIGUIENTES CLASIFICACIONES DE MATERIALES PELIGROSOS (INCLUYE CONTROL DE FUGAS, DERRAMES Y FUEGO CON LOS RESPECTIVOS MATERIALES)",
		"10-5-1": "EXPLOSIVOS",

		"10-5-2":  "GASES COMBUSTIBLES (DIFERENTES A GLP, GAS NATURAL, GAS DE CUIDAD Y MONÓXIDO DE CARBONO)",
		"10-5-3":  "GASES NO COMBUSTIBLES NI VENENOSOS",
		"10-5-4":  "GASES VENENOSOS",
		"10-5-5":  "LÍQUIDOS COMBUSTIBLES",
		"10-5-6":  "LÍQUIDOS INFLAMABLES",
		"10-5-7":  "SÓLIDOS COMBUSTIBLES",
		"10-5-8":  "SÓLIDOS QUE REACCIONAN CON AIRE",
		"10-5-9":  "SÓLIDOS QUE REACCIONAN CON AGUA",
		"10-5-10": "COMBURENTES",
		"10-5-11": "PERÓXIDOS ORGÁNICOS",
		"10-5-12": "VENENOS SÓLIDOS O LÍQUIDOS",
		"10-5-13": "AGENTES BIOLÓGICOS",
		"10-5-14": "RADIACTIVOS",
		"10-5-15": "CORROSIVOS",
		"10-5-16": "MISELÁNEOS O NO CLASIFICADOS",

		"10-6":    "ACCIDENTES CON GASES COMBUSTIBLES (GLP, GAS NATURAL, GAS DE CUIDAD Y MONÓXIDO DE CARBONO",
		"10-6-1":  "GAS LICUADO DE PETRÓLEO PROPANO BUTANO DESDE CILINDROS PORTATILES",
		"10-6-2":  "GAS LICUADO DE PETRÓLEO PROPANO BUTANO DESDE ESTANQUES FIJOS",
		"10-6-3":  "GAS LICUADO DE PETRÓLEO PROPANO BUTANO DESDE CAMIÓN A GRANEL",
		"10-6-4":  "GAS NATURAL DESDE UNA INSTALACIÓN DOMICLIARIA",
		"10-6-5":  "GAS NATURAL DESDE UNA INSTALACIÓN INDUSTRIAL",
		"10-6-6":  "GAS NATURAL DESDE CAÑERÍA EN VÍA PÚBLICA O GASEODUCTO",
		"10-6-7":  "GAS DE CIUDAD DESDE INSTALACIONES DOMICILIARIAS",
		"10-6-8":  "GAS DE CIUDAD DESDE INSTALACIONES INDUSTRIALES",
		"10-6-9":  "GAS DE CIUDAD DESDE CAÑERÍA O GAS DE CIUDAD EN VIA PÚBLICA",
		"10-6-10": "PRESENCIA DE MONÓXIDO CARBONO EN INTERIOR DE DOMICILIO",
		"10-6-11": "PRESENCIA DE MONÓXIDO CARBONO EN INTERIOR DE INDUSTRIA U OFICINA",
		"10-6-12": "EMERGENCIAS DE VEHÍCULOS PROPULSADOS A GASES COMBUSTIBLES",

		"10-6-13": "EMERGENCIA CON GASES DESDE ALCANTARILLAS DERIVADOS DEL PETRÓLEO U ORGÁNICOS",
		"10-6-14": "EXPLOSIÓN DE CILINDROS CON GASES COMPRIMIDOS",

		"10-6-15": "EXPLOSIÓN DE VAPORES EN EXPANSIÓN DE UN LÍQUIDO EN EBULLICIÓN (BLEVE)",

		"10-6-16": "EXPLOSIÓN POR COMBUSTIÓN (INDIVIDUALIZAR MATERIAL GENERADOR DE CALOR EN LA EXPLOSIÓN)",
		"10-7":    "Llamado eléctrico",
		"10-7-1":  "TENDIDO AÉREO DE ALTA TENSIÓN",
		"10-7-2":  "TENDIDO AÉREO DE BAJA TENSIÓN",
		"10-7-3":  "TENDIDO SUBTERRÁNEO DE ALTA TENSIÓN",
		"10-7-4":  "TENDIDO SUBTERRÁNEO DE BAJA TENSIÓN",
		"10-7-5":  "ACOMETIDA TRIFÁSICA",
		"10-7-6":  "ACOMETIDA MONOFÁSICA",
		"10-7-7":  "MEDIDOR DE ENERGÍA TRIFÁSICO",
		"10-7-8":  "MEDIDOR DE ENERGÍA MONOFÁSICO",
		"10-7-9":  "SUB- ESTACIÓN AÉREA (TRANSFORMADOR)",
		"10-8":    "Llamado no clasificado",
		"10-8-1":  "ABRIR CASA HABITACIÓN 1 PISO",
		"10-8-2":  "ABRIR CASA HABITACIÓN MÁS DE 1 PISO",
		"10-8-3":  "ABRIR DEPARTAMENTO U OFICINA DE 1º A 5º PISO",
		"10-8-4":  "ABRIR DEPARTAMENTO U OFICINA DE 6º A 10º PISO",
		"10-8-5":  "ABRIR DEPARTAMENTO U OFICINA DE 11º A 15º PISO",
		"10-8-6":  "ABRIR DEPARTAMENTO U OFICINA DE 16º A 20º PISO",
		"10-8-7":  "ABRIR DEPARTAMENTO U OFICINA DE 21º HACIA ARRIBA",
		"10-8-8":  "PROVISIÓN DE AGUA POTABLE",
		"10-8-9":  "PROVISIÓN DE ELECTRICIDAD CON GENERADORES",
		"10-8-10": "ASISTENCIA PREVENTIVA CON VEHICULO CONTRA INCENDIO",
		"10-8-11": "ASISTENCIA PREVENTIVA DE VEHÍCULO DE RESCATE",
		"10-8-12": "ASISTENCIA PREVENTIVA DE VEHICULO HAZMAT",
		"10-8-13": "TRABAJO EN ALTURA MAYOR A 10 MTS.",
		"10-8-14": "TRABAJO EN ALTURA MENOR A 10 MTS. (COLOCAR DRIZA O SIMILAR)",
		"10-8-15": "RASTREO DE PERSONAS EN SECTORES RURALES / URBANO",
		"10-8-16": "RASTREO DE PERSONAS BAJO EL AGUA (TRABAJO SUB-ACUÁTICO)",

		"10-8-17": "RASTREO DE PERSONAS SUPERFICIAL EN CAUCES DE AGUA O AGUAS ABIERTAS",
		"10-8-18": "RECUPERACIÓN DE CADÁVERES",

		"10-8-19": "INSTRUCCIÓN, ASESORÍAS, CAPACITACIÓN A TERCEROS O CAPACITACIONES INTERNAS",

		"10-8-20": "ACADEMIAS, GUARDIAS NOCTURNAS, DIURNAS, DESFILES, CAMPAÑAS ECONÓMICAS, REUNIONES",
		"10-8-21": "CAÍDA DE ÁRBOLES",
		"10-8-22": "VOLADURAS DE TECHO O TEMPORALES DE VIENTOS",
		"10-8-23": "EXTRACCIÓN DE AGUA POR INUNDACIÓN",
		"10-8-24": "RESCATE DE ANIMALES",
		"10-8-25": "EVACUACIÓN DEL CUERPO HACIA OTRO LUGAR",
		"10-8-26": "OPERATIVOS PLAN INTEGRAL DE SEGURIDAD ESCOLAR",
		"10-8-27": "DESPEJE DE CALLES",
		"10-8-28": "ATERRIZAJES DE HELICÓPTEROS",
		"10-8-29": "LAVADO DE CARRETERA (DERRAMES, ACCIDENTES, ETC.)",

		"10-8-30": "RESPUESTA PREVENTIVE A ATENTADOS (COORDINACIÓN COMANDO INCIDENTES)",
		"10-8-31": "RESPUESTA OPERATIVA A ATENTADOS (COORDINACIÓN COMANDO INCIDENTES)",
		"10-8-32": "CAÍDA DE AERONAVES EN ÁREAS URBANAS",
		"10-8-33": "CAÍDA DE AERONAVES EN ÁREAS RURALES",
		"10-8-34": "SALIDAS DE MÁQUINAS BOMBERILES",

		"10-9":   "SERVICIOS DE INSPECCIÓN Y REVISIÓN A DIVERSAS INSTALACIONES ECHAS POR BOMBEROS",
		"10-9-1": "PERITAJES A SOLICITUD DEL MINISTERIO PÚBLICO",
		"10-9-2": "REVISIÓN DE EDIFICIOS",
		"10-9-3": "INSPECCIÓN DE LOCAL COMERCIAL",
		"10-9-4": "INSPECCIÓN DE INDUSTRIAS",
		"10-9-5": "INSPECCIÓN DE GRIFOS",
		"10-9-6": "SIMULACROS O EJERCICIOS A TERCEROS",

		"10-10":   "CONCURRENCIA DE BOMBEROS A REBROTE DE FUEGO DE CUALQUIERA DE LAS ANTERIORES CATEGORÍAS",
		"10-10-0": "REBROTE DE FUEGOS DE CUALQUIER FUEGO ANTERIOR",
		"10-10-1": "REBROTE DE FUEGO EN VEHÍCULO",
		"10-10-2": "REBORTE DE FUEGO EN LUGAR ABIERTO",

		"10-11": "CONCURRENCIA DE BOMBEROS A APOYAR LABORES DE OTROS CUERPOS DE BOMBEROS",

		"10-11-1": "APOYO A OTROS CUERPOS (EN CUALQUIERA DE LOS ACTOS DEL SERVICIO DETALLADOS ANTERIORMENTE)",
	}
	return fireCodesMap
}
