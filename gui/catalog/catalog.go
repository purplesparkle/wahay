// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package gui

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"ar": &dictionary{index: arIndex, data: arData},
		"en": &dictionary{index: enIndex, data: enData},
		"es": &dictionary{index: esIndex, data: esData},
		"sv": &dictionary{index: svIndex, data: svData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%s\nMeeting ID: %s%%0D%%0A":          11,
	"A valid port is between 1 and 65535": 100,
	"Accept":                              30,
	"Allow the host to automatically join a newly created meeting": 31,
	"An error occurred\n\n%s":                                      16,
	"Are you sure you want to do this action?":                     32,
	"Are you sure you want to end this meeting?":                   33,
	"Are you sure you want to leave this meeting?":                 34,
	"Automatically join a meeting":                                 36,
	"Automatically join this meeting":                              35,
	"Automatically join this meeting when starting it":             37,
	"Be very careful. This information is sensitive and could potentially contain very private information. Only turn on these settings if you absolutely need it for debugging.": 38,
	"Browse": 39,
	"By clicking Yes, this meeting will end.":       40,
	"By clicking Yes, you will leave this meeting.": 41,
	"Cancel": 26,
	"Check this option to automatically join every meeting you host": 51,
	"Choose your email service to send invitation":                   52,
	"Client binary location":                                         42,
	"Configuration settings will be lost in the next session":        43,
	"Configure master password":                                      44,
	"Configured Mumble port is not valid: %s":                        3,
	"Confirmation":                   45,
	"Connecting, please wait...":     46,
	"Continue":                       47,
	"Copy Invitation":                48,
	"Copy Meeting ID":                49,
	"Copy URL":                       50,
	"Debugging":                      53,
	"Default Email":                  54,
	"Encrypt the configuration file": 55,
	"End this meeting":               57,
	"End this meeting for all":       58,
	"Error":                          0,
	"Finish":                         56,
	"General":                        59,
	"Gmail":                          60,
	"Host a new meeting":             61,
	"Host meeting":                   63,
	"Hosting":                        62,
	"If you backup the configuration file, we will reset the settings and continue normally. If the configuration file is encrypted, then we will ask you for a password to encrypt the new settings file.": 64,
	"If you disable this option, anyone could read your configuration settings":          24,
	"If you set this option to a file name, low level information will be logged there.": 65,
	"Invalid configuration file":                66,
	"Invalid password. Please, try again.":      67,
	"Invite others":                             68,
	"Join":                                      69,
	"Join Wahay Meeting":                        9,
	"Join a meeting":                            70,
	"Join meeting":                              71,
	"Join the meeting":                          72,
	"Join this meeting":                         73,
	"Keep configuration file when Wahay closes": 74,
	"Leave":              75,
	"Leave this meeting": 76,
	"Log debug info":     77,
	"Log debug output to the selected log file. If no file is selected then the log output will be written to the default log file.": 78,
	"Master password":                79,
	"Meeting ID":                     80,
	"Meeting ID:":                    81,
	"Meeting password":               82,
	"Mumble":                         83,
	"No, cancel":                     84,
	"Now you are hosting a meeting.": 85,
	"Open":                           27,
	"Open file":                      25,
	"Outlook":                        86,
	"Password":                       87,
	"Please enter the master password for the configuration file.":              88,
	"Please join the Wahay meeting with the following details:%%0D%%0A%%0D%%0A": 10,
	"Port":                               89,
	"Port out of range":                  90,
	"Raw log file":                       91,
	"Repeat the password":                92,
	"Save changes":                       93,
	"Security":                           94,
	"Settings":                           95,
	"Show":                               96,
	"Something went wrong: %s":           1,
	"Specify a password for the meeting": 97,
	"Start Meeting":                      13,
	"Start Meeting & Join":               12,
	"Start meeting":                      98,
	"The Meeting ID cannot be blank":     15,
	"The Mumble process is down":         14,
	"The error message":                  99,
	"The invitation email has been copied to the clipboard": 8,
	"The meeting ID has been copied to the clipboard":       7,
	"The meeting ID is invalid":                             18,
	"The meeting can't be closed: %s":                       4,
	"The onion service can't be deleted: %s":                6,
	"This action cannot be undone":                          101,
	"Toggle password visibility":                            102,
	"Type the Meeting ID (normally a .onion address)":       103,
	"Type the password":                                     104,
	"Type the password to join the meeting":                 105,
	"Type your preferred screen name":                       106,
	"Type your screen name":                                 107,
	"Username":                                              108,
	"Wahay is ready to use":                                 109,
	"We have detected that the configuration file is invalid or corrupted. Do you want to make a copy (backup) of it and continue?": 110,
	"We've found errors": 29,
	"Welcome":            111,
	"When this option is checked, the configuration settings will be stored in the device.": 112,
	"Yahoo Mail":                     113,
	"Yes, back it up &amp; continue": 114,
	"Yes, confirm":                   115,
	"You will not be asked for this password again until you restart Wahay.": 116,
	"enter a password at least 6 characters long":                            23,
	"enter the password confirmation":                                        21,
	"internal Tor instance has already been closed":                          5,
	"passwords do not match":                                                 22,
	"please enter a valid password":                                          20,
	"the Mumble client can not be used because: %s":                          19,
	"the provided meeting ID is invalid: \n\n%s":                             17,
	"tor can't be used":                                                      28,
	"we couldn't start the meeting":                                          2,
}

var arIndex = []uint32{ // 118 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000d, 0x0000001a, 0x00000027,
	0x00000034, 0x00000041, 0x0000004e, 0x0000005b,
	0x00000068, 0x00000075, 0x00000082, 0x0000008f,
	0x0000009c, 0x000000a9, 0x000000b6, 0x000000c3,
	0x000000d0, 0x000000dd, 0x000000ea, 0x000000f7,
	0x00000104, 0x00000111, 0x0000011e, 0x0000012b,
	0x00000138, 0x00000145, 0x00000152, 0x0000015f,
	0x0000016c, 0x00000179, 0x00000186, 0x00000193,
	// Entry 20 - 3F
	0x000001a0, 0x000001ad, 0x000001ba, 0x000001c7,
	0x000001d4, 0x000001e1, 0x000001ee, 0x000001fb,
	0x00000208, 0x00000215, 0x00000222, 0x0000022f,
	0x0000023c, 0x00000249, 0x00000256, 0x00000286,
	0x00000293, 0x000002a0, 0x000002ad, 0x000002ba,
	0x000002c7, 0x000002d4, 0x000002e1, 0x000002ee,
	0x000002fb, 0x00000308, 0x00000315, 0x00000322,
	0x0000032f, 0x0000033c, 0x00000349, 0x00000356,
	// Entry 40 - 5F
	0x00000363, 0x00000370, 0x0000037d, 0x0000038a,
	0x00000397, 0x000003a4, 0x000003b1, 0x000003be,
	0x000003cb, 0x000003d8, 0x000003e5, 0x000003f2,
	0x000003ff, 0x0000040c, 0x00000419, 0x00000426,
	0x00000433, 0x00000440, 0x0000044d, 0x0000045a,
	0x00000467, 0x00000474, 0x00000481, 0x0000048e,
	0x0000049b, 0x000004a8, 0x000004b5, 0x000004c2,
	0x000004cf, 0x000004dc, 0x000004e9, 0x000004f6,
	// Entry 60 - 7F
	0x00000503, 0x00000510, 0x0000051d, 0x0000052a,
	0x00000537, 0x00000544, 0x00000551, 0x0000055e,
	0x0000056b, 0x00000578, 0x00000585, 0x00000592,
	0x0000059f, 0x000005ac, 0x000005b9, 0x000005c6,
	0x000005d3, 0x000005e0, 0x000005ed, 0x000005fa,
	0x00000607, 0x00000614,
} // Size: 496 bytes

const arData string = "" + // Size: 1556 bytes
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02...الاتصال الرجاء الانتظار\x02TRANSLATE ME\x02TRANSL" +
	"ATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME"

var enIndex = []uint32{ // 118 elements
	// Entry 0 - 1F
	0x00000000, 0x00000006, 0x00000022, 0x00000040,
	0x0000006b, 0x0000008e, 0x000000bc, 0x000000e6,
	0x00000116, 0x0000014c, 0x0000015f, 0x000001a5,
	0x000001c3, 0x000001d8, 0x000001e6, 0x00000201,
	0x00000220, 0x00000239, 0x00000265, 0x0000027f,
	0x000002b0, 0x000002ce, 0x000002ee, 0x00000305,
	0x00000331, 0x0000037b, 0x00000385, 0x0000038c,
	0x00000391, 0x000003a3, 0x000003b6, 0x000003bd,
	// Entry 20 - 3F
	0x000003fa, 0x00000423, 0x0000044e, 0x0000047b,
	0x0000049b, 0x000004b8, 0x000004e9, 0x00000595,
	0x0000059c, 0x000005c4, 0x000005f2, 0x00000609,
	0x00000641, 0x0000065b, 0x00000668, 0x00000683,
	0x0000068c, 0x0000069c, 0x000006ac, 0x000006b5,
	0x000006f4, 0x00000721, 0x0000072b, 0x00000739,
	0x00000758, 0x0000075f, 0x00000770, 0x00000789,
	0x00000791, 0x00000797, 0x000007aa, 0x000007b2,
	// Entry 40 - 5F
	0x000007bf, 0x00000885, 0x000008d8, 0x000008f3,
	0x00000918, 0x00000926, 0x0000092b, 0x0000093a,
	0x00000947, 0x00000958, 0x0000096a, 0x00000994,
	0x0000099a, 0x000009ad, 0x000009bc, 0x00000a3b,
	0x00000a4b, 0x00000a56, 0x00000a62, 0x00000a73,
	0x00000a7a, 0x00000a85, 0x00000aa4, 0x00000aac,
	0x00000ab5, 0x00000af2, 0x00000af7, 0x00000b09,
	0x00000b16, 0x00000b2a, 0x00000b37, 0x00000b40,
	// Entry 60 - 7F
	0x00000b49, 0x00000b4e, 0x00000b71, 0x00000b7f,
	0x00000b91, 0x00000bb5, 0x00000bd2, 0x00000bed,
	0x00000c1d, 0x00000c2f, 0x00000c55, 0x00000c75,
	0x00000c8b, 0x00000c94, 0x00000caa, 0x00000d28,
	0x00000d30, 0x00000d86, 0x00000d91, 0x00000db0,
	0x00000dbd, 0x00000e04,
} // Size: 496 bytes

const enData string = "" + // Size: 3588 bytes
	"\x02Error\x02Something went wrong: %[1]s\x02we couldn't start the meetin" +
	"g\x02Configured Mumble port is not valid: %[1]s\x02The meeting can't be " +
	"closed: %[1]s\x02internal Tor instance has already been closed\x02The on" +
	"ion service can't be deleted: %[1]s\x02The meeting ID has been copied to" +
	" the clipboard\x02The invitation email has been copied to the clipboard" +
	"\x02Join Wahay Meeting\x02Please join the Wahay meeting with the followi" +
	"ng details:%0D%0A%0D%0A\x02%[1]s\x0aMeeting ID: %[2]s%0D%0A\x02Start Mee" +
	"ting & Join\x02Start Meeting\x02The Mumble process is down\x02The Meetin" +
	"g ID cannot be blank\x02An error occurred\x0a\x0a%[1]s\x02the provided m" +
	"eeting ID is invalid: \x0a\x0a%[1]s\x02The meeting ID is invalid\x02the " +
	"Mumble client can not be used because: %[1]s\x02please enter a valid pas" +
	"sword\x02enter the password confirmation\x02passwords do not match\x02en" +
	"ter a password at least 6 characters long\x02If you disable this option," +
	" anyone could read your configuration settings\x02Open file\x02Cancel" +
	"\x02Open\x02tor can't be used\x02We've found errors\x02Accept\x02Allow t" +
	"he host to automatically join a newly created meeting\x02Are you sure yo" +
	"u want to do this action?\x02Are you sure you want to end this meeting?" +
	"\x02Are you sure you want to leave this meeting?\x02Automatically join t" +
	"his meeting\x02Automatically join a meeting\x02Automatically join this m" +
	"eeting when starting it\x02Be very careful. This information is sensitiv" +
	"e and could potentially contain very private information. Only turn on t" +
	"hese settings if you absolutely need it for debugging.\x02Browse\x02By c" +
	"licking Yes, this meeting will end.\x02By clicking Yes, you will leave t" +
	"his meeting.\x02Client binary location\x02Configuration settings will be" +
	" lost in the next session\x02Configure master password\x02Confirmation" +
	"\x02Connecting, please wait...\x02Continue\x02Copy Invitation\x02Copy Me" +
	"eting ID\x02Copy URL\x02Check this option to automatically join every me" +
	"eting you host\x02Choose your email service to send invitation\x02Debugg" +
	"ing\x02Default Email\x02Encrypt the configuration file\x02Finish\x02End " +
	"this meeting\x02End this meeting for all\x02General\x02Gmail\x02Host a n" +
	"ew meeting\x02Hosting\x02Host meeting\x02If you backup the configuration" +
	" file, we will reset the settings and continue normally. If the configur" +
	"ation file is encrypted, then we will ask you for a password to encrypt " +
	"the new settings file.\x02If you set this option to a file name, low lev" +
	"el information will be logged there.\x02Invalid configuration file\x02In" +
	"valid password. Please, try again.\x02Invite others\x02Join\x02Join a me" +
	"eting\x02Join meeting\x02Join the meeting\x02Join this meeting\x02Keep c" +
	"onfiguration file when Wahay closes\x02Leave\x02Leave this meeting\x02Lo" +
	"g debug info\x02Log debug output to the selected log file. If no file is" +
	" selected then the log output will be written to the default log file." +
	"\x02Master password\x02Meeting ID\x02Meeting ID:\x02Meeting password\x02" +
	"Mumble\x02No, cancel\x02Now you are hosting a meeting.\x02Outlook\x02Pas" +
	"sword\x02Please enter the master password for the configuration file." +
	"\x02Port\x02Port out of range\x02Raw log file\x02Repeat the password\x02" +
	"Save changes\x02Security\x02Settings\x02Show\x02Specify a password for t" +
	"he meeting\x02Start meeting\x02The error message\x02A valid port is betw" +
	"een 1 and 65535\x02This action cannot be undone\x02Toggle password visib" +
	"ility\x02Type the Meeting ID (normally a .onion address)\x02Type the pas" +
	"sword\x02Type the password to join the meeting\x02Type your preferred sc" +
	"reen name\x02Type your screen name\x02Username\x02Wahay is ready to use" +
	"\x02We have detected that the configuration file is invalid or corrupted" +
	". Do you want to make a copy (backup) of it and continue?\x02Welcome\x02" +
	"When this option is checked, the configuration settings will be stored i" +
	"n the device.\x02Yahoo Mail\x02Yes, back it up &amp; continue\x02Yes, co" +
	"nfirm\x02You will not be asked for this password again until you restart" +
	" Wahay."

var esIndex = []uint32{ // 118 elements
	// Entry 0 - 1F
	0x00000000, 0x00000006, 0x0000001d, 0x0000003d,
	0x00000073, 0x00000098, 0x000000ae, 0x000000db,
	0x0000010f, 0x00000148, 0x00000171, 0x000001cc,
	0x000001f1, 0x0000020c, 0x0000021e, 0x0000023d,
	0x00000266, 0x0000027f, 0x000002b3, 0x000002d5,
	0x00000306, 0x00000334, 0x00000363, 0x00000381,
	0x000003b5, 0x00000413, 0x00000421, 0x0000042a,
	0x00000430, 0x00000445, 0x00000461, 0x00000469,
	// Entry 20 - 3F
	0x000004c1, 0x000004f3, 0x00000529, 0x0000055b,
	0x00000583, 0x000005b8, 0x000005ed, 0x000006a1,
	0x000006aa, 0x000006db, 0x0000070b, 0x00000729,
	0x0000076b, 0x0000078a, 0x00000798, 0x000007b8,
	0x000007c2, 0x000007d5, 0x000007eb, 0x000007f6,
	0x0000085c, 0x000008a1, 0x000008ad, 0x000008c2,
	0x000008e6, 0x000008ef, 0x00000905, 0x00000926,
	0x0000092e, 0x00000934, 0x00000951, 0x00000959,
	// Entry 40 - 5F
	0x0000096d, 0x00000a69, 0x00000ad0, 0x00000af4,
	0x00000b1f, 0x00000b2f, 0x00000b36, 0x00000b4c,
	0x00000b61, 0x00000b76, 0x00000b8d, 0x00000bc3,
	0x00000bc9, 0x00000be0, 0x00000c06, 0x00000cc4,
	0x00000cd8, 0x00000cea, 0x00000cfd, 0x00000d18,
	0x00000d1f, 0x00000d2c, 0x00000d53, 0x00000d5b,
	0x00000d67, 0x00000da9, 0x00000db0, 0x00000dc6,
	0x00000ddb, 0x00000df1, 0x00000e01, 0x00000e0b,
	// Entry 60 - 7F
	0x00000e1b, 0x00000e23, 0x00000e50, 0x00000e64,
	0x00000e78, 0x00000eab, 0x00000ecd, 0x00000ef1,
	0x00000f32, 0x00000f49, 0x00000f7a, 0x00000fa1,
	0x00000fbe, 0x00000fd0, 0x00000fee, 0x00001070,
	0x0000107b, 0x000010d0, 0x000010e0, 0x000010fd,
	0x0000110b, 0x00001154,
} // Size: 496 bytes

const esData string = "" + // Size: 4436 bytes
	"\x02Error\x02Algo salió mal: %[1]s\x02no se pudo comenzar la reunión\x02" +
	"El puerto configurado para Mumble es inválido: %[1]s\x02La reunión no se" +
	" pudo cerrar: %[1]s\x02tor no está definido\x02El servicio Onion no se p" +
	"udo eliminar: %[1]s\x02El ID de la reunion ha sido copiado al portapapel" +
	"es\x02El correo de invitación ha sido copiado al portapapeles\x02Unirse " +
	"a una reunión a travéz de Wahay\x02Por favor únete a la reunión a travéz" +
	" de Wahay con los siguientes detalles:%0D%0A%0D%0A\x02%[1]s\x0aID de la " +
	"reunión: %[2]s%0D%0A\x02Comenzar reunión y Unirse\x02Comenzar reunión" +
	"\x02El proceso Mumble está caído\x02El ID de la reunión no puede ser vac" +
	"ío\x02Ocurrió un error\x0a\x0a%[1]s\x02el ID de la reunión provisto es " +
	"inválido: \x0a\x0a%[1]s\x02El ID de la reunión es inválido\x02El cliente" +
	" Mumble no se puede usar porque: %[1]s\x02por favor especifique una cont" +
	"raseña válida\x02especifique la confirmación de la contraseña\x02las con" +
	"traseñas no coinciden\x02especifique una contraseña de mínimo 6 caracter" +
	"es\x02Si deshabilita esta opción, cualquier persona podría leer los pará" +
	"metros de configuración\x02Abrir archivo\x02Cancelar\x02Abrir\x02tor no " +
	"se puede usar\x02Encontramos algunos errores\x02Aceptar\x02Permita que e" +
	"l organizador se una automáticamente a una reunión cuando cree una nueva" +
	"\x02¿Está seguro de que quieres hacer esta acción?\x02¿Está seguro de qu" +
	"e quieres terminar esta reunión?\x02¿Está seguro de que quiere dejar est" +
	"a reunión?\x02Unirse automáticamente a esta reunión\x02Únase automáticam" +
	"ente a esta reunión al iniciarla\x02Únase automáticamente a esta reunión" +
	" al iniciarla\x02Ten mucho cuidado. Esta información es confidencial y p" +
	"odría contener información muy privada. Solo cambia esta configuración s" +
	"i la necesita absolutamente para la depuración.\x02Examinar\x02Al hacer " +
	"clic en Sí, esta reunión finalizará.\x02Al hacer clic en Sí, saldrá de e" +
	"sta reunión.\x02Ubicación del binario Mumble\x02Los ajustes de configura" +
	"ción se perderán en la próxima sesión\x02Configurar contraseña maestra" +
	"\x02Confirmación\x02Conectando, espere por favor...\x02Continuar\x02Copi" +
	"ar invitación\x02Copiar ID de reunión\x02Copiar URL\x02Marque esta opció" +
	"n para unirse automáticamente a cada reunión creada en la sección del an" +
	"fitrión\x02Elija su servicio de correo electrónico para enviar la invita" +
	"ción.\x02Depuración\x02Email predeterminado\x02Cifrar el archivo de conf" +
	"iguración\x02Terminar\x02Termina esta reunión\x02Termina esta reunión pa" +
	"ra todos\x02General\x02Gmail\x02Organizar una nueva reunión\x02Hosting" +
	"\x02Reunión de acogida\x02Si realiza una copia de seguridad del archivo " +
	"de configuración, restableceremos la configuración y continuaremos norma" +
	"lmente. Si el archivo de configuración está cifrado, le pediremos una co" +
	"ntraseña para cifrar el nuevo archivo de configuración.\x02Si establece " +
	"esta opción en un nombre de archivo, la información de bajo nivel se reg" +
	"istrará allí.\x02Archivo de configuración inválido\x02Contraseña invalid" +
	"a. Inténtalo de nuevo.\x02Invitar a otros\x02Unirse\x02Unirse a una reun" +
	"ión\x02Unirse a la reunión\x02Únete a la reunión\x02Únete a esta reunión" +
	"\x02Mantener el archivo de configuración al cerrar Wahay\x02Salir\x02Sal" +
	"ir de esta reunión\x02Registrar información de depuración\x02Registre la" +
	" salida de depuración en el archivo de registro seleccionado. Si no se s" +
	"elecciona ningún archivo, la salida del registro se escribirá en el arch" +
	"ivo de registro predeterminado.\x02Contraseña maestra\x02ID de la reunió" +
	"n\x02ID de la reunión:\x02Contraseña de la reunión\x02Mumble\x02No, canc" +
	"elar\x02Ahora estás organizando una reunión.\x02Outlook\x02Contraseña" +
	"\x02Ingrese la contraseña maestra para el archivo de configuración.\x02P" +
	"uerto\x02Puerto fuera de rango\x02Archivo de registros\x02Repita la cont" +
	"raseña\x02Guardar cambios\x02Seguridad\x02Configuraciones\x02Mostrar\x02" +
	"Especifique una contraseña para la reunión\x02Comience a reunirse\x02El " +
	"mensaje de error\x02El rango de puertos válidos está entre 1 y 65535\x02" +
	"Esta acción no se puede deshacer\x02Alternar visibilidad de contraseña" +
	"\x02Escriba la ID de la reunión (normalmente una dirección .onion)\x02Es" +
	"cribe la contraseña\x02Escriba la contraseña para unirse a la reunión" +
	"\x02Escriba su nombre de usuario preferido\x02Escriba su nombre de usuar" +
	"io\x02Nombre de usuario\x02Wahay está listo para usarse\x02Hemos detecta" +
	"do que el archivo de configuración no es válido o está dañado. ¿Desea ha" +
	"cer una copia de seguridad y continuar?\x02Bienvenido\x02Cuando esta opc" +
	"ión está marcada, la configuración se guardará en el dispositivo.\x02Cor" +
	"reo de Yahoo\x02Sí, respaldarlo y continuar\x02Si, confirmar\x02No se le" +
	" volverá a solicitar esta contraseña hasta que reinicie Wahay."

var svIndex = []uint32{ // 118 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000d, 0x0000001a, 0x00000027,
	0x00000034, 0x00000041, 0x0000004e, 0x0000005b,
	0x00000068, 0x00000075, 0x00000082, 0x0000008f,
	0x0000009c, 0x000000a9, 0x000000b6, 0x000000c3,
	0x000000d0, 0x000000dd, 0x000000ea, 0x000000f7,
	0x00000104, 0x00000111, 0x0000011e, 0x0000012b,
	0x00000138, 0x00000145, 0x00000152, 0x0000015f,
	0x0000016c, 0x00000179, 0x00000186, 0x00000193,
	// Entry 20 - 3F
	0x000001a0, 0x000001ad, 0x000001ba, 0x000001c7,
	0x000001d4, 0x000001e1, 0x000001ee, 0x000001fb,
	0x00000208, 0x00000215, 0x00000222, 0x0000022f,
	0x0000023c, 0x00000249, 0x00000256, 0x00000272,
	0x0000027f, 0x0000028c, 0x00000299, 0x000002a6,
	0x000002b3, 0x000002c0, 0x000002cd, 0x000002da,
	0x000002e7, 0x000002f4, 0x00000301, 0x0000030e,
	0x0000031b, 0x00000328, 0x00000335, 0x00000342,
	// Entry 40 - 5F
	0x0000034f, 0x0000035c, 0x00000369, 0x00000376,
	0x00000383, 0x00000390, 0x0000039d, 0x000003aa,
	0x000003b7, 0x000003c4, 0x000003d1, 0x000003de,
	0x000003eb, 0x000003f8, 0x00000405, 0x00000412,
	0x0000041f, 0x0000042c, 0x00000439, 0x00000446,
	0x00000453, 0x00000460, 0x0000046d, 0x0000047a,
	0x00000487, 0x00000494, 0x000004a1, 0x000004ae,
	0x000004bb, 0x000004c8, 0x000004d5, 0x000004e2,
	// Entry 60 - 7F
	0x000004ef, 0x000004fc, 0x00000509, 0x00000516,
	0x00000523, 0x00000530, 0x0000053d, 0x0000054a,
	0x00000557, 0x00000564, 0x00000571, 0x0000057e,
	0x0000058b, 0x00000598, 0x000005a5, 0x000005b2,
	0x000005bf, 0x000005cc, 0x000005d9, 0x000005e6,
	0x000005f3, 0x00000600,
} // Size: 496 bytes

const svData string = "" + // Size: 1536 bytes
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02Ansluter, var god vänta...\x02TRANSLATE ME\x02TRANSL" +
	"ATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME"

	// Total table size 13100 bytes (12KiB); checksum: B294E4A8