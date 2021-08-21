package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 输入：names = ["John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"], synonyms = ["(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"]
// 输出：["John(27)","Chris(36)"]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/baby-names-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type record struct {
	m   map[string]bool
	min string
}

func trulyMostPopular(names []string, synonyms []string) []string {
	m := make(map[string]string)
	for _, synonym := range synonyms {
		synonymList := strings.Split(synonym, ",")
		first := string([]byte(synonymList[0])[1:])
		second := string([]byte(synonymList[1])[0 : len(synonymList[1])-1])

		firstRoot := findRootName(m, first)
		secondRoot := findRootName(m, second)
		if firstRoot < secondRoot {
			m[secondRoot] = firstRoot
		} else {
			m[firstRoot] = secondRoot
		}
	}

	countMap := make(map[string]int64)
	for _, value := range names {
		nameSplit := strings.Split(value, "(")
		name := nameSplit[0]
		nameExpected := findRootName(m, name)
		bytes := []byte(nameSplit[1])
		c, _ := strconv.Atoi(string(bytes[0 : len(bytes)-1]))
		cc, ok := countMap[nameExpected]
		if !ok {
			countMap[nameExpected] = int64(c)
		} else {
			countMap[nameExpected] = int64(c) + cc
		}
	}

	r := make([]string, 0, len(countMap))
	for k, v := range countMap {
		r = append(r, fmt.Sprintf("%s(%d)", k, v))
	}

	return r
}

func findRootName(m map[string]string, name string) string {
	for true {
		next, ok := m[name]
		if !ok {
			return name
		}

		if next == name {
			return name
		}

		name = next
	}

	return ""
}

func trulyMostPopular1(names []string, synonyms []string) []string {
	m := make(map[string]*record)
	for _, synonym := range synonyms {
		synonym1 := strings.ReplaceAll(strings.ReplaceAll(synonym, "(", ""), ")", "")
		synonymList := strings.Split(synonym1, ",")
		first := synonymList[0]
		second := synonymList[1]

		m1, ok := m[first]
		if !ok {
			m2, ok1 := m[second]
			if !ok1 {
				subM := make(map[string]bool)
				subM[first] = true
				subM[second] = true

				r := &record{m: subM}
				if first < second {
					r.min = first
				} else {
					r.min = second
				}

				m[first] = r
				m[second] = r
			} else {
				m2.m[first] = true
				if first < m2.min {
					m2.min = first
				}
				m[first] = m2
			}
		} else {
			m2, ok1 := m[second]
			if !ok1 {
				m1.m[second] = true
				if second < m1.min {
					m1.min = second
				}
				m[second] = m1
			} else {
				// 都已经存在了
				// v1 -> m1
				// v2 -> m2
				// 需要合并m1和m2
				mmerge := make(map[string]bool)
				r := &record{m: mmerge}
				if m1.min <= m2.min {
					r.min = m1.min
				} else {
					r.min = m2.min
				}

				for v3, _ := range m1.m {
					mmerge[v3] = true
					m[v3] = r
				}

				for v3, _ := range m2.m {
					mmerge[v3] = true
					m[v3] = r
				}
			}
		}
	}

	countMap := make(map[string]int64)
	for _, value := range names {
		nameSplit := strings.Split(value, "(")
		name := nameSplit[0]
		r, ok := m[name]
		var nameExpected string
		if ok {
			nameExpected = r.min
		} else {
			nameExpected = name
		}
		bytes := []byte(nameSplit[1])
		c, _ := strconv.Atoi(string(bytes[0 : len(bytes)-1]))
		fmt.Printf("name:%v, expected:%v, count:%d\n", name, nameExpected, c)
		cc, ok := countMap[nameExpected]
		if !ok {
			countMap[nameExpected] = int64(c)
		} else {
			countMap[nameExpected] = int64(c) + cc
		}
	}

	r := make([]string, 0, len(countMap))
	for k, v := range countMap {
		r = append(r, fmt.Sprintf("%s(%d)", k, v))
	}

	return r
}

func main() {
	//names := []string{"John(15)", "Jon(12)", "Chris(13)", "Kris(4)", "Christopher(19)"}
	//synonyms := []string{"(Jon,John)", "(John,Johnny)", "(Chris,Kris)", "(Chris,Christopher)"}

	//names := []string{"Fcclu(70)", "Ommjh(63)", "Dnsay(60)", "Qbmk(45)", "Unsb(26)", "Gauuk(75)", "Wzyyim(34)", "Bnea(55)", "Kri(71)", "Qnaakk(76)", "Gnplfi(68)", "Hfp(97)", "Qoi(70)", "Ijveol(46)", "Iidh(64)", "Qiy(26)", "Mcnef(59)", "Hvueqc(91)", "Obcbxb(54)", "Dhe(79)", "Jfq(26)", "Uwjsu(41)", "Wfmspz(39)", "Ebov(96)", "Ofl(72)", "Uvkdpn(71)", "Avcp(41)", "Msyr(9)", "Pgfpma(95)", "Vbp(89)", "Koaak(53)", "Qyqifg(85)", "Dwayf(97)", "Oltadg(95)", "Mwwvj(70)", "Uxf(74)", "Qvjp(6)", "Grqrg(81)", "Naf(3)", "Xjjol(62)", "Ibink(32)", "Qxabri(41)", "Ucqh(51)", "Mtz(72)", "Aeax(82)", "Kxutz(5)", "Qweye(15)", "Ard(82)", "Chycnm(4)", "Hcvcgc(97)", "Knpuq(61)", "Yeekgc(11)", "Ntfr(70)", "Lucf(62)", "Uhsg(23)", "Csh(39)", "Txixz(87)", "Kgabb(80)", "Weusps(79)", "Nuq(61)", "Drzsnw(87)", "Xxmsn(98)", "Onnev(77)", "Owh(64)", "Fpaf(46)", "Hvia(6)", "Kufa(95)", "Chhmx(66)", "Avmzs(39)", "Okwuq(96)", "Hrschk(30)", "Ffwni(67)", "Wpagta(25)", "Npilye(14)", "Axwtno(57)", "Qxkjt(31)", "Dwifi(51)", "Kasgmw(95)", "Vgxj(11)", "Nsgbth(26)", "Nzaz(51)", "Owk(87)", "Yjc(94)", "Hljt(21)", "Jvqg(47)", "Alrksy(69)", "Tlv(95)", "Acohsf(86)", "Qejo(60)", "Gbclj(20)", "Nekuam(17)", "Meutux(64)", "Tuvzkd(85)", "Fvkhz(98)", "Rngl(12)", "Gbkq(77)", "Uzgx(65)", "Ghc(15)", "Qsc(48)", "Siv(47)"}
	//synonyms := []string{"(Gnplfi,Qxabri)", "(Uzgx,Siv)", "(Bnea,Lucf)", "(Qnaakk,Msyr)", "(Grqrg,Gbclj)", "(Uhsg,Qejo)", "(Csh,Wpagta)", "(Xjjol,Lucf)", "(Qoi,Obcbxb)", "(Npilye,Vgxj)", "(Aeax,Ghc)", "(Txixz,Ffwni)", "(Qweye,Qsc)", "(Kri,Tuvzkd)", "(Ommjh,Vbp)", "(Pgfpma,Xxmsn)", "(Uhsg,Csh)", "(Qvjp,Kxutz)", "(Qxkjt,Tlv)", "(Wfmspz,Owk)", "(Dwayf,Chycnm)", "(Iidh,Qvjp)", "(Dnsay,Rngl)", "(Qweye,Tlv)", "(Wzyyim,Kxutz)", "(Hvueqc,Qejo)", "(Tlv,Ghc)", "(Hvia,Fvkhz)", "(Msyr,Owk)", "(Hrschk,Hljt)", "(Owh,Gbclj)", "(Dwifi,Uzgx)", "(Iidh,Fpaf)", "(Iidh,Meutux)", "(Txixz,Ghc)", "(Gbclj,Qsc)", "(Kgabb,Tuvzkd)", "(Uwjsu,Grqrg)", "(Vbp,Dwayf)", "(Xxmsn,Chhmx)", "(Uxf,Uzgx)"}

	names := []string{"Pwsuo(71)", "Prf(48)", "Rgbu(49)", "Zvzm(31)", "Xxcl(25)", "Bbcpth(42)", "Padz(70)", "Jmqqsj(19)", "Uwy(26)", "Jylbla(65)", "Xioal(11)", "Npbu(62)", "Jpftyg(96)", "Tal(46)", "Hnc(100)", "Yldu(85)", "Alqw(45)", "Wbcxi(34)", "Kxjw(36)", "Clplqf(8)", "Fayxe(66)", "Slfwyo(48)", "Xbesji(70)", "Pmbz(22)", "Oip(2)", "Fzoe(63)", "Qync(79)", "Utc(11)", "Sqwejn(19)", "Ngi(8)", "Gsiiyo(60)", "Bcs(73)", "Icsvku(1)", "Yzwm(92)", "Vaakt(21)", "Uvt(70)", "Axaqkm(100)", "Gyhh(84)", "Gaoo(98)", "Ghlj(35)", "Umt(13)", "Nfimij(52)", "Zmeop(77)", "Vje(29)", "Rqa(47)", "Upn(89)", "Zhc(44)", "Slh(66)", "Orpqim(69)", "Vxs(85)", "Gql(19)", "Sfjdjc(62)", "Ccqunq(93)", "Oyo(32)", "Bvnkk(52)", "Pxzfjg(45)", "Kaaht(28)", "Arrugl(57)", "Vqnjg(50)", "Dbufek(63)", "Fshi(62)", "Lvaaz(63)", "Phlto(41)", "Lnow(70)", "Mqgga(31)", "Adlue(82)", "Zqiqe(27)", "Mgs(46)", "Zboes(56)", "Dma(70)", "Jnij(57)", "Ghk(14)", "Mrqlne(39)", "Ljkzhs(35)", "Rmlbnj(42)", "Qszsny(93)", "Aasipa(26)", "Wzt(41)", "Xuzubb(90)", "Maeb(56)", "Mlo(18)", "Rttg(4)", "Kmrev(31)", "Kqjl(39)", "Iggrg(47)", "Mork(88)", "Lwyfn(50)", "Lcp(42)", "Zpm(5)", "Qlvglt(36)", "Liyd(48)", "Jxv(67)", "Xaq(70)", "Tkbn(81)", "Rgd(85)", "Ttj(28)", "Ndc(62)", "Bjfkzo(54)", "Lqrmqh(50)", "Vhdmab(41)"}
	synonyms := []string{"(Uvt,Rqa)", "(Qync,Kqjl)", "(Fayxe,Upn)", "(Maeb,Xaq)", "(Pmbz,Vje)", "(Hnc,Dma)", "(Pwsuo,Gyhh)", "(Gyhh,Aasipa)", "(Fzoe,Lcp)", "(Mgs,Vhdmab)", "(Qync,Rgd)", "(Gql,Liyd)", "(Gyhh,Tkbn)", "(Arrugl,Adlue)", "(Wbcxi,Slfwyo)", "(Yzwm,Vqnjg)", "(Lnow,Vhdmab)", "(Lvaaz,Rttg)", "(Nfimij,Iggrg)", "(Vje,Lqrmqh)", "(Jylbla,Ljkzhs)", "(Jnij,Mlo)", "(Adlue,Zqiqe)", "(Qync,Rttg)", "(Gsiiyo,Vxs)", "(Xxcl,Fzoe)", "(Dbufek,Xaq)", "(Ccqunq,Qszsny)", "(Zmeop,Mork)", "(Qync,Ngi)", "(Zboes,Rmlbnj)", "(Yldu,Jxv)", "(Padz,Gsiiyo)", "(Oip,Utc)", "(Tal,Pxzfjg)", "(Adlue,Zpm)", "(Bbcpth,Mork)", "(Qync,Lvaaz)", "(Pmbz,Qync)", "(Alqw,Ngi)", "(Bcs,Maeb)", "(Rgbu,Zmeop)"}

	fmt.Printf("%s", trulyMostPopular(names, synonyms))
}
