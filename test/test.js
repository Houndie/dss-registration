const assert = require('chai').assert;
const puppeteer = require('puppeteer');
describe('Registration', function() {
	this.timeout(40000);
	it('should register successfully', async function() {
		const browser = await puppeteer.launch();
		const page = await browser.newPage();

		page.on('console', msg => {
			if (msg.type() != 'error' && msg.type() != 'warning') {
				return;
			}
			for (let i = 0; i < msg.args().length; ++i) {
				assert.fail(`${i}: ${msg.args()[i]}`);
			}
		});

		await page.goto('http://localhost:8081/registration/');

		await page.waitForSelector('form');

		await page.type('input[name=firstName]', 'James');
		await page.type('input[name=lastName]', 'Benze');
		await page.type('input[name=streetAddress]', '123 any street');
		await page.type('input[name=city]', 'Gotham');
		await page.select('select[name=state]', 'AZ');
		await page.type('input[name=zipCode]', '12345');
		await page.type('input[name=email]', 'benzejaa@gmail.com');
		await page.type('input[name=homeScene]', 'gotham knights');
		await page.click('input[name=isStudent]');
		await page.select('select[name=passType]', 'Full');
		await page.select('select[name=level]', '0');
		await page.click('input[name=mixAndMatch]');
		await page.select('select[name=role]', '0');
		await page.click('input[name=soloJazz]');
		await page.click('input[name=teamCompetition]');
		await page.type('input[name=teamName]', 'The big biggertons');
		await page.click('input[name=tshirt]');
		await page.select('select[name=style]', "0");
		await page.select('select[name=housing]', "Provide");
		await page.type('input[name=pets]', 'I have a dog');
		await page.type('input[name=quantity]', "1");
		await page.type('textarea[name=provideDetails]', "all friends are welcome");

		await Promise.all([
			page.waitForNavigation(),
			page.click('button[type=submit]')
		])

		assert.equal(page.url().split('?')[0], "https://connect.squareupsandbox.com/v2/checkout")

		return browser.close();
	});
});
