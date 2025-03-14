/*
  Warnings:

  - You are about to drop the column `checkNumber` on the `Bank` table. All the data in the column will be lost.
  - You are about to drop the column `nominals` on the `Bank` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[phoneNumber]` on the table `Bank` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `accountNumber` to the `Bank` table without a default value. This is not possible if the table is not empty.
  - Added the required column `nominal` to the `Bank` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Bank" DROP COLUMN "checkNumber",
DROP COLUMN "nominals",
ADD COLUMN     "accountNumber" TEXT NOT NULL,
ADD COLUMN     "nominal" DOUBLE PRECISION NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX "Bank_phoneNumber_key" ON "Bank"("phoneNumber");
