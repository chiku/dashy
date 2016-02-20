var Stage = require("../../src/components/Stage");

describe("Stage", function() {
    describe("#render", function() {
        var stage = new Stage().render({
            "name": "Test",
            "status": "building"
        });

        it("creates a DOM representation", function() {
            expect(stage[0]).toEqual("div");
        });

        it("has CSS class based on its status", function() {
            expect(stage[1]).toEqual({
                "class": "stage building"
            });
        });

        it("has contents based on its name", function() {
            expect(stage[2]).toEqual("Test");
        });

        describe("when status is not in all lower-case", function() {
            var stage = new Stage().render({
                "name": "Test",
                "status": "Building"
            });

            it("has a lower-name CSS class name", function() {
                expect(stage[1]).toEqual({
                    "class": "stage building"
                });
            });
        });
    });
});
