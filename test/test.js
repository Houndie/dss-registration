const assert = require('chai').assert;
const puppeteer = require('puppeteer');

let browser
let page

before(async () => {
	browser = await puppeteer.launch({
		args: ['--no-sandbox', '--disable-setuid-sandbox']
	})
	page = await browser.newPage()
})

registrations = [
	{
		testName: 'registration 1',
		firstName: 'James',
		lastName: 'Benze',
		streetAddress: '123 any street',
		city: 'gotham',
		state: 'arizona',
		zipCode: '12345',
		email: 'benzejaa@gmail.com',
		homeScene: 'gotham knights',
		isStudent: true,
		passType: 'Full',
		level: '0',
		mixAndMatch: true,
		role: '0',
		soloJazz: true,
		teamCompetition: true,
		teamName: 'The big biggertons',
		tshirt: true,
		style: '0',
		housing: 'Provide',
		pets: 'I have a dog',
		quantity: '1',
		details: 'all friends are welcome'
	}
]

describe('Registration', function() {
	this.timeout(40000);
	registrations.forEach((r) => {
		it('should register successfully with '+r.testName, async function() {
			if (typeof process.env.FRONTEND === 'undefined') {
				assert.fail('FRONTEND environment variable not defined')
			}
			const frontend = process.env.FRONTEND

			page.on('console', msg => {
				if (msg.type() != 'error' && msg.type() != 'warning') {
					return;
				}
				for (let i = 0; i < msg.args().length; ++i) {
					console.log(`${i}: ${msg.args()[i]}`);
				}
			});

			await page.goto(frontend+'/registration/');

			await page.waitForSelector('form');

			await page.type('input[name=firstName]', r.firstName);
			await page.type('input[name=lastName]', r.lastName);
			await page.type('input[name=streetAddress]', r.streetAddress);
			await page.type('input[name=city]', r.city);
			await page.select('select[name=state]', r.state);
			await page.type('input[name=zipCode]', r.zipCode);
			await page.type('input[name=email]', r.email);
			await page.type('input[name=homeScene]', r.homeScene);
			if (r.isStudent) {
				await page.click('input[name=isStudent]');
			}
			await page.select('select[name=passType]', r.passType);
			if (r.passType == 'Full') {
				await page.select('select[name=level]', '0');
			}
			if (r.mixAndMatch) {
				await page.click('input[name=mixAndMatch]');
				await page.select('select[name=role]', r.role);
			}
			if (r.soloJazz) {
				await page.click('input[name=soloJazz]');
			}
			if (r.teamCompetition) {
				await page.click('input[name=teamCompetition]');
				await page.type('input[name=teamName]', r.teamName);
			}
			if (r.tshirt) {
				await page.click('input[name=tshirt]');
				await page.select('select[name=style]', "0");
			}
			await page.select('select[name=housing]', r.housing);
			if (r.housing == 'Provide') {
				await page.type('input[name=pets]', 'I have a dog');
				await page.type('input[name=quantity]', "1");
				await page.type('textarea[name=provideDetails]', "all friends are welcome");
			}

			await Promise.all([
				page.waitForNavigation(),
				page.click('button[type=submit]')
			])

			assert.equal(page.url().split('?')[0], "https://connect.squareupsandbox.com/v2/checkout")
		});
	});
});

after(async () => {
	await browser.close()
})
